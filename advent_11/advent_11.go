package main

import "log"

const realFilename = "advent_11.test.txt"

func main() {
	inputFile := readAdvent11File(realFilename)
	result := calcAdvent11Result(inputFile)
	log.Printf("Answer: %d", result.answer)
}

type advent11Result struct {
	answer int
}

func calcAdvent11Result(inputFile advent11File) advent11Result {
	result := advent11Result{}
	return result
}

type advent11File struct {
	linesAmount int
	lines       []advent11FileLine
}

type advent11FileLine struct {
	screamersAmt    int
	screamerNumbers []int
}

func readAdvent11File(filename string) advent11File {
	result := advent11File{}
	return result
}
