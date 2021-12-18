package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
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
func lexicographicallyMinimalPermutations(variants []string) []string {
	perLen := make(map[int][]string, 0)
	for _, v := range variants {
		l := len(v)
		if _, ok := perLen[l]; !ok {
			perLen[l] = make([]string, 0)
		}
		perLen[l] = append(perLen[l], v)
	}
	result := make([]string, 0)
	for _, v := range perLen {
		minimalString := lexicographicallyMinimalString(v)
		result = append(result, minimalString)
	}

	return result
}
func shortestRuneArrays(variants ...[]rune) [][]rune {
	vs := make([]string, 0, len(variants))
	for _, v := range variants {
		vs = append(vs, string(v))
	}

	min := lexicographicallyMinimalPermutations(vs)

	res := make([][]rune, 0, len(min))
	for _, m := range min {
		res = append(res, bytes.Runes([]byte(m)))
	}

	return res

	//minLength := math.MaxInt
	//var shortArrays [][]rune
	//for _, v := range variants {
	//	curLen := len(v)
	//	if curLen < minLength {
	//		minLength = curLen
	//		shortArrays = make([][]rune, 0, 1)
	//		shortArrays = append(shortArrays, v)
	//	} else if curLen == minLength {
	//		shortArrays = append(shortArrays, v)
	//	}
	//}
	//return shortArrays
}

func asRunes(in string) []rune {
	return bytes.Runes([]byte(in))
}
func addPermutations(current [][]rune, r rune, amt int) [][]rune {
	variants := createPermutations(r, amt)
	result := make([][]rune, 0, len(current))

	// to save from OOM, only shortest variants will stay
	variants = shortestRuneArrays(variants...)
	for _, v := range variants {
		if len(current) == 0 {
			n := make([]rune, 0)
			n = append(n, v...)
			result = append(result, n)
		} else {
			for _, c := range current {
				n := make([]rune, 0)
				n = append(n, c...)
				n = append(n, v...)
				result = append(result, n)
			}
		}
	}
	return result
}
func createPermutations(r rune, amt int) [][]rune {
	var symbols []rune

	switch r {
	case '2':
		symbols = asRunes("ABC")
	case '3':
		symbols = asRunes("DEF")
	case '4':
		symbols = asRunes("GHI")
	case '5':
		symbols = asRunes("JKL")
	case '6':
		symbols = asRunes("MNO")
	case '7':
		symbols = asRunes("PQRS")
	case '8':
		symbols = asRunes("TUV")
	case '9':
		symbols = asRunes("WXYZ")
	}
	amounts := make([]int, len(symbols))
	leftover := amt
	for revIdx := len(symbols); revIdx > 0; revIdx-- {
		l := leftover
		leftover = l % revIdx
		amounts[revIdx-1] = l / revIdx
	}
	result := ""
	for idx, amount := range amounts {
		result += strings.Repeat(string(symbols[idx]), amount)
	}
	return append(make([][]rune, 0), asRunes(result))
}

func createCodeVariants(in string) []string {
	result := make([][]rune, 0)
	inLen := len(in)
	for runeIdx := 0; runeIdx < inLen; {

		r := in[runeIdx]

		// 1. need to find max available length of equal numbers
		addendum := 1 // amount of equal numbers
		for ; runeIdx+addendum < inLen; addendum++ {
			if in[runeIdx+addendum] != r {
				break
			}
		}

		result = addPermutations(result, rune(r), addendum)

		runeIdx += addendum
	}
	res := make([]string, 0, len(result))
	for _, r := range result {
		res = append(res, string(r))
	}
	return res
}

func calcAdvent8Result(inputFile advent8File) advent8Result {
	variants := createCodeVariants(string(inputFile))
	result := advent8Result{}
	if len(variants) != 0 {
		result.value = lexicographicallyMinimalString(variants)
	}
	return result
}

type advent8File string

func readAdvent8File(filename string) advent8File {
	readFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open file %q: %w", filename, err))
	}
	return advent8File(readFile)
}
