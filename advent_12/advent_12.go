package main

import (
	"log"
	"math"
)

func main() {
	result := calcAdvent12Result(14)
	log.Printf("Answer: %d (%v)", result.answer, result.multiples)
}

func calcAdvent12Result(digitsAmount int) advent12Result {
	return advent12Result{}
}

type advent12Result struct {
	answer    int64
	multiples [2]int64
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
