package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	maxInt32 := math.MaxInt32
	maxInt32Str := fmt.Sprintf("%d", maxInt32)
	log.Println(fmt.Sprintf("maxInt32 %d has %d characters", maxInt32, len(maxInt32Str)))
	maxInt64 := math.MaxInt64
	maxInt64Str := fmt.Sprintf("%d", maxInt64)
	log.Println(fmt.Sprintf("maxInt64 %d has %d characters", maxInt64, len(maxInt64Str)))
}

func newDigitsAmountChecker(digitsAmount int) digitsAmountChecker {
	return digitsAmountChecker{digitsAmount: digitsAmount,
		maxV: int64(math.Pow(10, float64(digitsAmount))) - 1,
		minV: int64(math.Pow(10, float64(digitsAmount-1)))}
}

type digitsAmountChecker struct {
	digitsAmount int
	maxV         int64
	minV         int64
}

func (dc digitsAmountChecker) isFit(v int64) bool {
	return v >= dc.minV && v <= dc.maxV
}
