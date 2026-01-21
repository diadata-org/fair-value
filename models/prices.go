package models

import (
	"encoding/json"
	"errors"
	"math"
	"math/big"
	"time"

	luminametacontract "github.com/diadata-org/fair-value/contracts/lumina/metacontract"
	"github.com/diadata-org/fair-value/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

// AssetQuotation is the most recent price point information on an asset.
type AssetQuotation struct {
	Asset  Asset     `json:"Asset"`
	Price  float64   `json:"Price"`
	Source string    `json:"Source"`
	Time   time.Time `json:"Time"`
}

// GetPrice returns the latest price of asset @a as published by DIA metacontract with @address.
// Fallback in case of error or 0 price is diadata API.
func (a *Asset) GetPrice(
	metacontractAddress common.Address,
	precision int,
	client *ethclient.Client,
) (assetQuotation AssetQuotation, err error) {

	assetQuotation, err = a.GetOnchainPrice(metacontractAddress, precision, client)
	if err != nil || assetQuotation.Price == 0 {
		log.Warnf("On-chain price not available for symbol-address-blockchain:%s -- %s -- %s", a.Symbol, a.Blockchain, a.Address)
		return a.GetPriceFromDiaAPI()
	}

	// TO DO: Should we add a threshold such that onchain price is only taken into account if timestamp not older?
	assetQuotation.Source = "DIA_Lumina_" + metacontractAddress.Hex()

	return
}

// GetOnchainPrice returns a quotation for asset @a by querying the metacontract with @metacontractAddress.
func (a *Asset) GetOnchainPrice(
	metacontractAddress common.Address,
	precision int,
	client *ethclient.Client,
) (aq AssetQuotation, err error) {

	var caller *luminametacontract.MetacontractCaller
	caller, err = luminametacontract.NewMetacontractCaller(metacontractAddress, client)
	if err != nil {
		return
	}

	priceBig, timeUnixBig, err := caller.GetValue(&bind.CallOpts{}, a.Symbol+"/USD")
	if err != nil {
		return
	}

	if priceBig == nil {
		err = errors.New("price is zero")
		return
	}

	aq.Price, _ = new(big.Float).Mul(new(big.Float).SetInt(priceBig), big.NewFloat(math.Pow10(-precision))).Float64()

	if timeUnixBig != nil {
		aq.Time = time.Unix(timeUnixBig.Int64(), 0)
	} else {
		aq.Time = time.Now()
	}

	aq.Asset = *a
	aq.Source = "DIA_Lumina_" + metacontractAddress.Hex()

	return
}

// GetPriceFromDiaAPI returns the price of the asset given by @address and @blockchain as returned by
// assetQuotation endpoint of diadata API.
func (a *Asset) GetPriceFromDiaAPI() (aq AssetQuotation, err error) {

	baseString := "https://api.diadata.org/v1/assetQuotation/" + a.Blockchain + "/" + a.Address

	var response []byte
	response, _, err = utils.GetRequest(baseString)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &aq)

	return
}

// TO DO: Update taking into account the new data types.

func ComputeContractExchangeRatePrice(totalUnderlying *big.Int, totalShares *big.Int) (price float64, err error) {
	return utils.DivideBigInts(totalUnderlying, totalShares)
}

func ComputeReserveBackingPrice(reserves *big.Int, supply *big.Int) (price float64, err error) {
	return utils.DivideBigInts(reserves, supply)
}

func ComputeNetAssetValuePrice(assets *big.Int, liabilities *big.Int, totalSupply *big.Int) (price float64, err error) {
	return utils.DivideBigInts(big.NewInt(0).Sub(assets, liabilities), totalSupply)
}

func ComputeRedemptionValuePrice(amountRedeemed *big.Int, amountIn *big.Int) (price float64, err error) {
	return utils.DivideBigInts(amountRedeemed, amountIn)
}
