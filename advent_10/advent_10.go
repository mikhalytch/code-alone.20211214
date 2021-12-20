package main

import "log"

func main() {
	filename := "advent_10.test.txt"
	inputFile := readAdvent10File(filename)
	result := calcAdvent10Result(inputFile)
	log.Printf("Answer: %d\n", result.answer)
}

type advent10Result struct {
	answer int
}

func calcAdvent10Result(inputFile advent10File) advent10Result {
	result := advent10Result{}
	return result
}

type advent10File struct {
	amount int
	sizes  []int
}

func readAdvent10File(filename string) advent10File {
	result := advent10File{}
	return result
}
