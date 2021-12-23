package main

import "log"

const realFilename = "advent_13.test.txt"

func main() {
	inputFile := readAdvent13File(realFilename)
	result := calcAdvent13Result(inputFile)
	log.Printf("Answer: %s", result.answer)
}

type advent13Result struct {
	answer string
}

func calcAdvent13Result(inputFile advent13File) advent13Result {
	result := advent13Result{}
	return result
}

type advent13File struct {
	amount, ascNum int
	phones         []string
}

func readAdvent13File(filename string) advent13File {
	result := advent13File{}
	return result
}
