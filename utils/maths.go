package utils

import (
	"errors"
	"math"
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

func Min(n, m uint64) uint64 {
	if n <= m {
		return n
	}
	return m
}

func ScaleFloat(f float64, decimals int) *big.Int {
	fBig := big.NewFloat(f)
	scaling := big.NewFloat(math.Pow10(decimals))
	priceScaled := new(big.Float).Mul(fBig, scaling)
	valueUSDInt := new(big.Int)
	priceScaled.Int(valueUSDInt)
	return valueUSDInt
}

func MulFloatAndIntToInt(x float64, y *big.Int) *big.Int {
	if y == nil {
		return big.NewInt(0)
	}
	xf := new(big.Float).SetPrec(256).SetFloat64(x)
	yf := new(big.Float).SetPrec(256).SetInt(y)
	zf := new(big.Float).Mul(xf, yf)

	out := new(big.Int)
	zf.Int(out) // truncates toward zero
	return out
}
