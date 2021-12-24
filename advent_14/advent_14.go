package main

import "log"

const realFilename = "advent_14.test.txt"

func main() {
	inputFile := readAdvent14File(realFilename)
	result := calcAdvent14Result(inputFile)
	log.Printf("Result: %q\n", result.answer)
}

type advent14Result struct {
	answer string
}

func calcAdvent14Result(inputFile advent14File) advent14Result {
	return advent14Result{}
}

type simpleRestriction struct {
	codePosition     int
	restrictedSymbol rune
}

type advent14File struct {
	codeLength, rulesAmount int
	rules                   []simpleRestriction
}

func readAdvent14File(filename string) advent14File {
	return advent14File{}
}

func calcMedian(set []rune) rune {
	return 0
}

func createLimitedSet(restricted []rune) []rune {
	return nil
}
