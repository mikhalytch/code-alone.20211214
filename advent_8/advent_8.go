package main

import (
	"log"
	"math"
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
	sort.Strings(shortStrings)
	return shortStrings[0]
}

func calcAdvent8Result(inputFile advent8File) advent8Result {
	return advent8Result{}
}

type advent8File struct {
}

func readAdvent8File(filename string) advent8File {
	return advent8File{}
}
