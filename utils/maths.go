package utils

import (
	"errors"
	"math/big"
)

func DivideBigInts(numerator *big.Int, denominator *big.Int) (result float64, err error) {
	if denominator.Cmp(big.NewInt(0)) == 0 {
		err = errors.New("totalShares must not be 0")
		return
	}
	result, _ = new(big.Float).Quo(big.NewFloat(0).SetInt(numerator), big.NewFloat(0).SetInt(denominator)).Float64()
	return
}
