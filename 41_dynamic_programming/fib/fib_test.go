package fib

import (
	"PixelAlg/common"
	"log"
	"testing"
)

func TestFib(t *testing.T) {
	defer common.TimeCost("Fib")
	n := 50
	log.Printf("Fib %v = %d", n, Fib(n))
}

func TestFibWithTmpRet(t *testing.T) {
	defer common.TimeCost("FibWithTmp")
	n := 50
	log.Printf("FibWithTmp %v = %d", n, FibWithTmpRet(n))
}
