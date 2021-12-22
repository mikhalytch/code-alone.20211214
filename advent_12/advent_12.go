package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
)

func main() {
	result := calcAdvent12Result(14)
	log.Printf("Answer: %d (%v)", result.answer, result.multiples)
}

type advent12Result struct {
	answer    int64
	multiples [2]int64
}

func (ar *advent12Result) addPrimes(p1 int64, p2 int64, dc *digitsAmountChecker, pc primesChecker) {
	m := p1 * p2
	if ar.answer < m && pc.isFit(p1, p2) && dc.isFit(m) {
		ar.answer = m
		ar.multiples = [2]int64{p1, p2}
	}
}

func calcAdvent12Result(digitsAmount int) advent12Result {
	digitsChecker := newDigitsAmountChecker(digitsAmount)
	pChecker := primesChecker(100)

	result := advent12Result{}
	currentPrime := int64(2)
	for digitsChecker.isBelowHighBorder(currentPrime) {
		variants := primeVariants(currentPrime, pChecker)
		ps := pairs(variants)
		addVariants(ps, &result, digitsChecker, pChecker)
		c, err := nextPrime(currentPrime, pChecker)
		if err != nil {
			break
		}
		currentPrime = c
	}

	return result
}

var noPrime = fmt.Errorf("no prime within reach")

// return `noPrime` error if no prime for checker limits exists
var b = new(big.Int)

func nextPrime(curPrime int64, checker primesChecker) (int64, error) {
	for n := curPrime + 1; checker.isFit(curPrime, n); n++ {
		b.SetInt64(n)
		if b.ProbablyPrime(0) {
			return n, nil
		}
	}
	return 0, noPrime
}
func primeVariants(curPrime int64, checker primesChecker) []int64 {
	result := make([]int64, 0, checker)
	next := curPrime
	for {
		if !checker.isFit(curPrime, next) {
			break
		}
		result = append(result, next)
		n, err := nextPrime(next, checker)
		if err != nil {
			break
		}
		next = n

	}
	return result
}
func pairs(variants []int64) [][2]int64 {
	res := make([][2]int64, 0, len(variants))
	v0 := variants[0]
	for _, v := range variants {
		res = append(res, [2]int64{v0, v})
	}
	return res
}
func addVariants(variants [][2]int64, result *advent12Result, checker *digitsAmountChecker, checker2 primesChecker) {
	for _, v := range variants {
		result.addPrimes(v[0], v[1], checker, checker2)
	}
}

func newDigitsAmountChecker(digitsAmount int) *digitsAmountChecker {
	return &digitsAmountChecker{digitsAmount: digitsAmount,
		maxV: int64(math.Pow(10, float64(digitsAmount))) - 1,
		minV: int64(math.Pow(10, float64(digitsAmount-1)))}
}

type digitsAmountChecker struct {
	digitsAmount int
	maxV         int64
	minV         int64
}

func (dc digitsAmountChecker) isFit(v int64) bool {
	return dc.isAboveLowBorder(v) && dc.isBelowHighBorder(v)
}

func (dc digitsAmountChecker) isBelowHighBorder(v int64) bool {
	return v <= dc.maxV
}

func (dc digitsAmountChecker) isAboveLowBorder(v int64) bool {
	return v >= dc.minV
}

type primesChecker int

func (pc primesChecker) isFit(p1, p2 int64) bool {
	return int64Abs(p1-p2) <= int64(pc)
}
func int64Abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
