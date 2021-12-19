package main

import "fmt"

func main() {
	filename := "advent_9.test.txt"
	inputFile := readAdvent9File(filename)
	result := calcAdvent9Result(inputFile)
	fmt.Printf("Answer is: %v", result.answer)
}

type advent9Result struct {
	answer int
}

func calcAdvent9Result(inputFile advent9File) advent9Result {
	return advent9Result{}
}

type advent9File struct {
	amount  int
	numbers []int
}

func readAdvent9File(filename string) advent9File {
	return advent9File{}
}
