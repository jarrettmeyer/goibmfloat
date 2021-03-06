package goibmfloat

import (
	"math"
)

func Float64(bytes []byte) float64 {
	sign := bytes[0] >> 7
	posOrNeg := float64(1 - 2*int(sign))
	exponent := bytes[0] & 0x7f
	fraction := 0.0
	for i := 0; i < 24; i++ {
		numerator := getBit(bytes, 8+i)
		denominator := float64(int(2 << uint(i)))
		fraction += numerator / denominator
	}
	value := posOrNeg * math.Pow(16.0, float64(exponent-64)) * fraction
	return value
}

func getBit(bytes []byte, n int) float64 {
	index := int(n / 8)
	shift := 7 - math.Mod(float64(n), 8.0)
	return float64(bytes[index] >> uint(shift) & 1)
}
