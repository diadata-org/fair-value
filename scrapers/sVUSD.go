package scrapers

import (
	"fmt"
	"math/big"

	"github.com/diadata-org/fair-value/models"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	erc20 "github.com/diadata-org/fair-value/contracts/erc20"
	svusd "github.com/diadata-org/fair-value/contracts/vetro/svusd"
)

type SVUSDScraper struct {
	BaseScraper
	client         *ethclient.Client
	config         models.FeedConfig
	blockchain     string
	metacontractFV models.MetacontractData
}

func NewSVUSDScraper(
	config models.FeedConfig,
	metacontractData models.MetacontractData,
	metacontractFV models.MetacontractData,
) (*SVUSDScraper, error) {

	client, err := ethclient.Dial(utils.Getenv("RPC_NODE_SVUSD", ""))
	if err != nil {
		return nil, fmt.Errorf("sVUSD -- make eth client for %s: %v", config.Symbol, err)
	}

	scraper := SVUSDScraper{
		BaseScraper:    NewBaseScraper(metacontractData),
		client:         client,
		config:         config,
		blockchain:     config.Blockchain,
		metacontractFV: metacontractFV,
	}
	return &scraper, nil
}

func (s *SVUSDScraper) TotalShares() (*big.Int, error) {

	r, err := svusd.NewSvusdCaller(common.HexToAddress(s.config.Address), s.client)
	if err != nil {
		return nil, err
	}

	return r.TotalSupply(&bind.CallOpts{})
}

func (s *SVUSDScraper) TotalUnderlying() (totalUnderlying *big.Int, totalValueUnderlying *big.Int, err error) {
	r, err := svusd.NewSvusdCaller(common.HexToAddress(s.config.Address), s.client)
	if err != nil {
		return
	}

	totalUnderlying, err = r.TotalAssets(&bind.CallOpts{})
	if err != nil {
		return
	}

	// Compute USD Value of totalUnderlying using underlying asset from contract (VUSD)
	underlyingAsset, err := s.GetUnderlyingAsset(r)
	if err != nil {
		return
	}
	// First try to fetch price from fair-value metacontract.
	quoteUnderlying, err := underlyingAsset.GetPrice(s.metacontractFV.Address, s.metacontractFV.Precision, s.metacontractFV.Client)
	if err != nil {
		quoteUnderlying, err = underlyingAsset.GetPrice(s.metacontractData.Address, s.metacontractData.Precision, s.metacontractData.Client)
		if err != nil {
			return
		}
	}

	totalValueUnderlying = utils.MulFloatAndIntToInt(quoteUnderlying.Price, totalUnderlying)

	return
}

func (s *SVUSDScraper) GetUnderlyingAsset(r *svusd.SvusdCaller) (asset models.Asset, err error) {

	assetAddress, err := r.Asset(&bind.CallOpts{})
	if err != nil {
		return
	}
	erc20Caller, err := erc20.NewERC20MetadataCaller(assetAddress, s.client)
	if err != nil {
		return
	}

	symbol, err := erc20Caller.Symbol(&bind.CallOpts{})
	if err != nil {
		return
	}
	asset.Blockchain = s.blockchain
	asset.Address = assetAddress.Hex()
	asset.Symbol = symbol

	return

}

func (s *SVUSDScraper) DataChannel() chan models.FairValueData { return s.dataChannel }
func (s *SVUSDScraper) GetConfig() models.FeedConfig           { return s.config }
func (s *SVUSDScraper) Close() chan bool                       { return s.BaseScraper.Close() }
