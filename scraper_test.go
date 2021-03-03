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

// func TestGetLimitPrice(t *testing.T) {
// 	output := GetLimitPrice()
// 	if output != 2.18 {
// 		t.Error("Test Failed: {}")
// 	}
// }

func BenchMarkGetCoins(b *testing.B) {
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
