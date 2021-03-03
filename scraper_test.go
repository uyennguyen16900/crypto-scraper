package main

import (
	"testing"
)

func TestGetCoins(t *testing.T) {
	coins := []Crypto{
		{"WashHigh", "$51.342", "GWHS"},
		{"Marina", "$5.6", "MA"},
		{"Polkadot", "$37.68", "DOT"},
		{"XRP", "$0.4449", "XRP"},
		{"Litecoin", "$193.66", "LTC"},
		{"Chainlink", "$30.58", "LINK"},
		{"MK", "$34,567.234", "MK"},
		{"UN", "$16.2", "UN"},
	}

	expect := []string{"MA", "DOT", "XRP", "LINK", "UN"}
	output := GetCoins(51.1, coins)

	if len(expect) != len(output) {
		t.Error("Test failed! Expected: {}, received: {}", expect, output)
	}
	for i, v := range expect {
		if v != output[i] {
			t.Error("Test failed! Expected: {}, received: {}", expect, output)
		}
	}
}

func TestConvertStringPriceToFloat(t *testing.T) {
	tests := []struct {
		input  string
		expect float64
	}{
		{"$45,124.67", 45124.67},
		{"$0.4263", 0.4263},
		{"$0.00001223", 0.00001223},
		{"$879", 879},
	}

	for _, test := range tests {
		if output := ConvertStringPriceToFloat(test.input); output != test.expect {
			t.Error("Test failed! Expected: {}, received: {}", test.expect, output)
		}
	}
}

func TestSortPrice(t *testing.T) {
	coins := []Crypto{
		{"WashHigh", "$51.342", "GWHS"},
		{"Marina", "$5.6", "MA"},
		{"Polkadot", "$37.68", "DOT"},
		{"XRP", "$0.4449", "XRP"},
		{"Litecoin", "$193.66", "LTC"},
		{"Chainlink", "$30.58", "LINK"},
		{"MK", "$34,567.234", "MK"},
		{"UN", "$16.2", "UN"},
	}

	expect := []Crypto{
		{"XRP", "$0.4449", "XRP"},
		{"Marina", "$5.6", "MA"},
		{"UN", "$16.2", "UN"},
		{"Chainlink", "$30.58", "LINK"},
		{"Polkadot", "$37.68", "DOT"},
		{"WashHigh", "$51.342", "GWHS"},
		{"Litecoin", "$193.66", "LTC"},
		{"MK", "$34,567.234", "MK"},
	}

	output := SortPrice(coins)

	for i, v := range expect {
		if v != output[i] {
			t.Error("Test failed! Expected: {}, received: {}", expect, output)
		}
	}
}

// func TestGetLimitPrice(t *testing.T) {
// 	output := GetLimitPrice()
// 	if output != 2.18 {
// 		t.Error("Test Failed: {}")
// 	}
// }

func BenchmarkGetCoins(b *testing.B) {
	coins := []Crypto{
		{"Polkadot", "$37.68", "DOT"},
		{"XRP", "$0.4449", "XRP"},
		{"Litecoin", "$193.66", "LTC"},
		{"Chainlink", "$30.58", "LINK"},
	}
	for i := 0; i < b.N; i++ {
		GetCoins(20.87, coins)
	}
}
