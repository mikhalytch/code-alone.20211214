package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	filename := "advent_8.test.txt"
	inputFile := readAdvent8File(filename)
	result := calcAdvent8Result(inputFile)
	log.Printf("Result: %s", result.value)
}

type advent8Result struct {
	value string
}

func lexicographicallyMinimalString(variants []string) string {
	shortStrings := shortestStrings(variants)
	sort.Strings(shortStrings)
	return shortStrings[0]
}

func shortestStrings(variants []string) []string {
	minLength := math.MaxInt
	var shortStrings []string
	for _, v := range variants {
		curLen := len(v)
		if curLen < minLength {
			minLength = curLen
			shortStrings = make([]string, 0, 1)
			shortStrings = append(shortStrings, v)
		} else if curLen == minLength {
			shortStrings = append(shortStrings, v)
		}
	}
	return shortStrings
}
func shortestRuneArrays(variants ...[]rune) [][]rune {
	minLength := math.MaxInt
	var shortArrays [][]rune
	for _, v := range variants {
		curLen := len(v)
		if curLen < minLength {
			minLength = curLen
			shortArrays = make([][]rune, 0, 1)
			shortArrays = append(shortArrays, v)
		} else if curLen == minLength {
			shortArrays = append(shortArrays, v)
		}
	}
	return shortArrays
}

func asRunes(in string) []rune {
	return bytes.Runes([]byte(in))
}
func addPermutations(current [][]rune, variants ...[]rune) [][]rune {
	result := make([][]rune, 0, len(current))
	if len(current) == 0 {
		current = append(current, make([]rune, 0))
	}
	for _, c := range current {
		// to save from OOM, only shortest variants will stay
		shortVariants := shortestRuneArrays(variants...)
		for _, v := range shortVariants {
			n := make([]rune, len(c))
			copy(n, c)
			n = append(n, v...)
			result = append(result, n)
		}
	}
	return result
}

func createCodeVariants(in string) []string {
	result := make([][]rune, 0)
	inLen := len(in)
	for runeIdx := 0; runeIdx < inLen; {
		r := in[runeIdx]
		r2b := false
		if runeIdx+1 < inLen && in[runeIdx+1] == r {
			r2b = true
		}
		r3b := false
		if r2b && runeIdx+2 < inLen && in[runeIdx+2] == r {
			r3b = true
		}
		r4b := false
		if r3b && runeIdx+3 < inLen && in[runeIdx+3] == r {
			r4b = true
		}
		switch r {
		case '2': // 3
			if r3b {
				result = addPermutations(result, asRunes("AAA"), asRunes("AB"), asRunes("BA"), asRunes("C"))
				runeIdx += 3
			} else if r2b {
				result = addPermutations(result, asRunes("AA"), asRunes("B"))
				runeIdx += 2
			} else {
				result = addPermutations(result, asRunes("A"))
				runeIdx += 1
			}
		case '3': // 3
			if r3b {
				result = addPermutations(result, asRunes("DDD"), asRunes("DE"), asRunes("ED"), asRunes("F"))
				runeIdx += 3
			} else if r2b {
				result = addPermutations(result, asRunes("DD"), asRunes("E"))
				runeIdx += 2
			} else {
				result = addPermutations(result, asRunes("D"))
				runeIdx += 1
			}
		case '4': // 3
			if r3b {
				result = addPermutations(result, asRunes("GGG"), asRunes("GH"), asRunes("HG"), asRunes("I"))
				runeIdx += 3
			} else if r2b {
				result = addPermutations(result, asRunes("GG"), asRunes("H"))
				runeIdx += 2
			} else {
				result = addPermutations(result, asRunes("G"))
				runeIdx += 1
			}
		case '5': // 3
			if r3b {
				result = addPermutations(result, asRunes("JJJ"), asRunes("JK"), asRunes("KJ"), asRunes("L"))
				runeIdx += 3
			} else if r2b {
				result = addPermutations(result, asRunes("JJ"), asRunes("K"))
				runeIdx += 2
			} else {
				result = addPermutations(result, asRunes("J"))
				runeIdx += 1
			}
		case '6': // 3
			if r3b {
				result = addPermutations(result, asRunes("MMM"), asRunes("MN"), asRunes("NM"), asRunes("O"))
				runeIdx += 3
			} else if r2b {
				result = addPermutations(result, asRunes("MM"), asRunes("N"))
				runeIdx += 2
			} else {
				result = addPermutations(result, asRunes("M"))
				runeIdx += 1
			}
		case '7': // 4
			if r4b {
				result = addPermutations(result, asRunes("PPPP"), asRunes("QPP"), asRunes("PQP"), asRunes("PPQ"), asRunes("QQ"), asRunes("PR"), asRunes("RP"), asRunes("S"))
				runeIdx += 4
			} else if r3b {
				result = addPermutations(result, asRunes("PPP"), asRunes("PQ"), asRunes("QP"), asRunes("R"))
				runeIdx += 3
			} else if r2b {
				result = addPermutations(result, asRunes("PP"), asRunes("Q"))
				runeIdx += 2
			} else {
				result = addPermutations(result, asRunes("P"))
				runeIdx += 1
			}
		case '8': // 3
			if r3b {
				result = addPermutations(result, asRunes("TTT"), asRunes("TU"), asRunes("UT"), asRunes("V"))
				runeIdx += 3
			} else if r2b {
				result = addPermutations(result, asRunes("TT"), asRunes("U"))
				runeIdx += 2
			} else {
				result = addPermutations(result, asRunes("T"))
				runeIdx += 1
			}
		case '9': // 4
			if r4b {
				result = addPermutations(result, asRunes("WWWW"), asRunes("XWW"), asRunes("WWX"), asRunes("WXW"), asRunes("XX"), asRunes("WY"), asRunes("YW"), asRunes("Z"))
				runeIdx += 4
			} else if r3b {
				result = addPermutations(result, asRunes("WWW"), asRunes("WX"), asRunes("XW"), asRunes("Y"))
				runeIdx += 3
			} else if r2b {
				result = addPermutations(result, asRunes("WW"), asRunes("X"))
				runeIdx += 2
			} else {
				result = addPermutations(result, asRunes("W"))
				runeIdx += 1
			}
		default:
			log.Fatalf("unexpected %v at index %d", r, runeIdx)
		}
	}
	res := make([]string, 0, len(result))
	for _, r := range result {
		res = append(res, string(r))
	}
	return res
}

func calcAdvent8Result(inputFile advent8File) advent8Result {
	variants := createCodeVariants(string(inputFile))
	min := lexicographicallyMinimalString(variants)
	return advent8Result{min}
}

type advent8File string

func readAdvent8File(filename string) advent8File {
	readFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open file %q: %w", filename, err))
	}
	return advent8File(readFile)
}
