package goibmfloat

import (
  "math"
	"testing"
)

const (
  tolerance = 1e-6
)

type readTest struct {
	bytes  []byte
	number float64
}

var golden = []readTest{
	{[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0},
	{[]byte{0x42, 0x80, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00}, 128.50},
	{[]byte{0xc2, 0x80, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00}, -128.50},
	{[]byte{0x42, 0x7b, 0x73, 0x33, 0x33, 0x33, 0x33, 0x34}, 123.449997},
}

func TestToFloat64(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		result, err := ToFloat64(g.bytes)
		if err != nil {
			t.Fatalf("%s", err)
		}
		if math.Abs(result - g.number) > tolerance {
			t.Fatalf("ToFloat64: want %f, got %f.", g.number, result)
		}
	}
}
