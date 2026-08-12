package scrapers

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	chainlinkaggregatorv3 "github.com/diadata-org/fair-value/contracts/chainlink/aggregatorv3"
	erc20 "github.com/diadata-org/fair-value/contracts/erc20"
	ierc4626 "github.com/diadata-org/fair-value/contracts/ierc4626"
	vetrogateway "github.com/diadata-org/fair-value/contracts/vetro/gateway"
	vetrotreasury "github.com/diadata-org/fair-value/contracts/vetro/treasury"
	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ----------------------------------------------------------------------------
// VUSD fair-value scraper
//
// Model (basket-backed stablecoin):
//   reserve_usd   = Σ_i (buffer_i + deployed_i) * oracle_i_usd
//   backed_supply = total_supply - amo_supply
//   fair_value    = reserve_usd / backed_supply
//   value_usd     = min(1, fair_value)
//
// Feeder fields (1e18):
//   numerator   = reserve_usd
//   denominator = backed_supply
//   fairValue   = fair_value
//   valueUsd    = value_usd
//
// NOTE on TotalUnderlying():
// VUSD is multi-collateral; there is no single natural "amount underlying" token quantity.
// For framework compatibility we return totalSupply as placeholder `totalUnderlying`, while
// `totalValueUnderlying` is the economically meaningful reserve USD value.
// ----------------------------------------------------------------------------

var (
	oneE18Int = big.NewInt(1_000_000_000_000_000_000)
	zeroAddr  = common.Address{}
)

type VUSDScraper struct {
	BaseScraper
	client *ethclient.Client

	config       models.FeedConfig
	blockchain   string
	treasuryAddr common.Address
}

type vusdResult struct {
	TotalUnderlying1e18 *big.Int // sum of collateral amounts (normalized to 1e18)
	totalValueUSD       *big.Int // USD value of totalUnderlying
	totalSupply         *big.Int // minted VUSD tokens
	ValueUsd1e18        *big.Int // totalValueUSD / totalSupply
	Flags               []string
}

func NewVUSDScraper(config models.FeedConfig, metacontractData models.MetacontractData) *VUSDScraper {

	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_VUSD", ""))
	if err != nil {
		log.Errorf("VUSD -- make eth client for %s: %v", config.Symbol, err)
		return nil
	}

	return &VUSDScraper{
		BaseScraper:  NewBaseScraper(metacontractData),
		client:       client,
		config:       config,
		blockchain:   config.Blockchain,
		treasuryAddr: common.HexToAddress(config.Address),
	}
}

func (s *VUSDScraper) TotalShares() (*big.Int, error) {
	r, err := s.compute()
	if err != nil {
		return nil, err
	}

	return r.totalSupply, nil
}

func (s *VUSDScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {
	result, err := s.compute()
	if err != nil {
		return nil, nil, err
	}

	return result.TotalUnderlying1e18, result.totalValueUSD, nil
}

func (s *VUSDScraper) compute() (vusdResult, error) {

	treasury, err := vetrotreasury.NewVetrotreasuryCaller(s.treasuryAddr, s.client)
	if err != nil {
		return vusdResult{}, fmt.Errorf("treasury caller: %w", err)
	}

	// get peggedTokenAddress from treasury contract. Should be VUSD address.
	peggedTokenAddress, err := treasury.PEGGEDTOKEN(&bind.CallOpts{})
	if err != nil {
		return vusdResult{}, fmt.Errorf("call PEGGEDTOKEN: %w", err)
	}
	vusd, err := erc20.NewERC20Caller(peggedTokenAddress, s.client)
	if err != nil {
		return vusdResult{}, fmt.Errorf("vusd erc20 caller: %w", err)
	}

	// Get Gateway address from treasury contract.
	gatewayAddress, err := treasury.Gateway(&bind.CallOpts{})
	if err != nil {
		return vusdResult{}, fmt.Errorf("call Gateway address: %w", err)
	}

	// Get total supply. Subtract AMO if flag is set to true in config file. Defaults to false.
	totalSupplyRaw, err := vusd.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return vusdResult{}, fmt.Errorf("vusd totalSupply: %w", err)
	}
	if totalSupplyRaw.Sign() <= 0 {
		return vusdResult{}, fmt.Errorf("non-positive totalSupply")
	}

	subtractAmo := s.config.Params[1].([]any)
	if len(s.config.Params) == 2 && len(subtractAmo) > 0 && subtractAmo[0].(string) == "true" {
		gateway, err := vetrogateway.NewVetrogatewayCaller(gatewayAddress, s.client)
		if err != nil {
			return vusdResult{}, fmt.Errorf("gateway caller: %w", err)
		}

		amoSupplyRaw, err := gateway.AmoSupply(&bind.CallOpts{})
		if err != nil {
			return vusdResult{}, fmt.Errorf("gateway amoSupply: %w", err)
		}
		totalSupplyRaw = big.NewInt(0).Sub(totalSupplyRaw, amoSupplyRaw)
	}

	// Get underlying tokens and their total USD value.
	tokens, err := treasury.WhitelistedTokens(&bind.CallOpts{})
	if err != nil {
		return vusdResult{}, fmt.Errorf("treasury whitelistedTokens: %w", err)
	}
	if len(tokens) == 0 {
		return vusdResult{}, fmt.Errorf("no whitelisted tokens")
	}

	flags := make([]string, 0, 8)
	totalUnderlying1e18 := big.NewInt(0)
	reserveUSD := new(big.Float).SetPrec(256).SetFloat64(0)

	for _, token := range tokens {
		cfg, err := treasury.TokenConfig(&bind.CallOpts{}, token)
		if err != nil {
			return vusdResult{}, fmt.Errorf("tokenConfig(%s): %w", token.Hex(), err)
		}

		tokenCaller, err := erc20.NewERC20MetadataCaller(token, s.client)
		if err != nil {
			return vusdResult{}, fmt.Errorf("erc20 caller(%s): %w", token.Hex(), err)
		}
		tokenDecU8, err := tokenCaller.Decimals(&bind.CallOpts{})
		if err != nil {
			return vusdResult{}, fmt.Errorf("token decimals(%s): %w", token.Hex(), err)
		}
		tokenDec := int(tokenDecU8)

		// buffer balance
		bufferRaw, err := tokenCaller.BalanceOf(&bind.CallOpts{}, s.treasuryAddr)
		if err != nil {
			return vusdResult{}, fmt.Errorf("buffer balanceOf(%s): %w", token.Hex(), err)
		}

		// deployed balance via vault
		deployedRaw := big.NewInt(0)
		if !sameAddress(cfg.Vault, zeroAddr) {
			vault, err := ierc4626.NewIerc4626Caller(cfg.Vault, s.client)
			if err != nil {
				return vusdResult{}, fmt.Errorf("vault caller(%s): %w", cfg.Vault.Hex(), err)
			}
			sharesRaw, err := vault.BalanceOf(&bind.CallOpts{}, s.treasuryAddr)
			if err != nil {
				return vusdResult{}, fmt.Errorf("vault balanceOf(%s): %w", cfg.Vault.Hex(), err)
			}
			if sharesRaw.Sign() > 0 {
				deployedRaw, err = vault.ConvertToAssets(&bind.CallOpts{}, sharesRaw)
				if err != nil {
					return vusdResult{}, fmt.Errorf("vault convertToAssets(%s): %w", cfg.Vault.Hex(), err)
				}
			}
		}

		amountRaw := new(big.Int).Add(bufferRaw, deployedRaw)
		if amountRaw.Sign() == 0 {
			continue
		}

		// totalUnderlying: aggregated collateral amount (normalized)
		totalUnderlying1e18.Add(totalUnderlying1e18, normalizeTo1e18(amountRaw, tokenDec))

		// totalValueUnderlying leg: Chainlink USD price
		oracle, err := chainlinkaggregatorv3.NewChainlinkaggregatorv3Caller(cfg.Oracle, s.client)
		if err != nil {
			return vusdResult{}, fmt.Errorf("oracle caller(%s): %w", cfg.Oracle.Hex(), err)
		}
		rd, err := oracle.LatestRoundData(&bind.CallOpts{})
		if err != nil {
			return vusdResult{}, fmt.Errorf("latestRoundData(%s): %w", cfg.Oracle.Hex(), err)
		}
		oracleDecU8, err := oracle.Decimals(&bind.CallOpts{})
		if err != nil {
			return vusdResult{}, fmt.Errorf("oracle decimals(%s): %w", cfg.Oracle.Hex(), err)
		}
		oracleDec := int(oracleDecU8)

		if rd.Answer.Sign() <= 0 {
			flags = append(flags, "non_positive_oracle:"+cfg.Oracle.Hex())
			continue
		}

		// USD leg = (amount / 10^tokenDec) * (price / 10^oracleDec)
		leg := mulBigFloat(intToFloat(amountRaw), scaleDownInt(rd.Answer, oracleDec))
		leg = scaleDownFloat(leg, tokenDec)
		reserveUSD = new(big.Float).Add(reserveUSD, leg)
	}

	return vusdResult{
		TotalUnderlying1e18: totalUnderlying1e18,
		totalValueUSD:       floatTo1e18(reserveUSD),
		totalSupply:         new(big.Int).Set(totalSupplyRaw),
		Flags:               flags,
	}, nil
}

func (s *VUSDScraper) DataChannel() chan models.FairValueData { return s.dataChannel }
func (s *VUSDScraper) GetConfig() models.FeedConfig           { return s.config }
func (s *VUSDScraper) Close() chan bool                       { return s.BaseScraper.Close() }

// -----------------------------------------------------------------------------
// Helpers
// -----------------------------------------------------------------------------

func sameAddress(a, b common.Address) bool {
	return strings.EqualFold(a.Hex(), b.Hex())
}

func intToFloat(x *big.Int) *big.Float {
	return new(big.Float).SetPrec(256).SetInt(x)
}

func scaleDownInt(x *big.Int, decimals int) *big.Float {
	if decimals <= 0 {
		return intToFloat(x)
	}
	den := new(big.Float).SetPrec(256).SetFloat64(math.Pow10(decimals))
	return new(big.Float).Quo(intToFloat(x), den)
}

func scaleDownFloat(x *big.Float, decimals int) *big.Float {
	if decimals <= 0 {
		return new(big.Float).SetPrec(256).Set(x)
	}
	den := new(big.Float).SetPrec(256).SetFloat64(math.Pow10(decimals))
	return new(big.Float).Quo(x, den)
}

func mulBigFloat(a, b *big.Float) *big.Float {
	return new(big.Float).Mul(a, b)
}

func floatTo1e18(x *big.Float) *big.Int {
	if x == nil {
		return big.NewInt(0)
	}
	scaled := new(big.Float).Mul(x, new(big.Float).SetInt(oneE18Int))
	out := new(big.Int)
	scaled.Int(out) // truncate toward zero, matches Python int()
	return out
}

func normalizeTo1e18(amount *big.Int, tokenDec int) *big.Int {
	out := new(big.Int).Set(amount)
	switch {
	case tokenDec == 18:
		return out
	case tokenDec < 18:
		m := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18-tokenDec)), nil)
		return out.Mul(out, m)
	default:
		d := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDec-18)), nil)
		return out.Div(out, d)
	}
}
