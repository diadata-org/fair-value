package models

import (
	"math/big"

	"github.com/diadata-org/fair-value/utils"
)

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
