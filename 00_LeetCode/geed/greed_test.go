package geed

import (
	"fmt"
	"testing"
)

var (
	COIN = []int{1, 2, 5}
	//COIN = []int{2}
)

func TestJustPrint(t *testing.T) {
	//coin := FindMinCoin(COIN, 1)
	//fmt.Println(coin)
	coin := FindMinCoin(COIN, 3)
	fmt.Println(coin)
	coin = FindMinCoin(COIN, 6)
	fmt.Println(coin)
}
