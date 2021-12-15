package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "advent_2.sample"
	input := readAdvent2File(filename)
	result := calcAdvent2(input)
	log.Printf("Answer: %d", result.time)
}

type advent2File struct {
	width, point int
}

func readAdvent2File(filename string) advent2File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening file %q: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("error closing file %q: %w", filename, err))
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	result := advent2File{}

	for index := 0; scanner.Scan(); index++ {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(fmt.Errorf("error converting number at index %d: %w", index, err))
		}

		switch index {
		case 0:
			result.width = number
		case 1:
			result.point = number
		default:
			log.Fatalf("read unknown number %d at index %d", number, index)
		}
	}

	return result
}

type advent2Result struct {
	time int
}

func calcAdvent2(input advent2File) advent2Result {
	cycle1 := 2 * (input.point /*- 0*/)
	cycle2 := 2 * (input.width - input.point)
	return advent2Result{lcm(cycle1, cycle2)}
}

func lcm(a, b int) int {
	ab := a * b
	lcm := ab / gcd(a, b)
	return lcm
}

/*
from wikipedia: https://en.wikipedia.org/wiki/Extended_Euclidean_algorithm#Pseudocode
function extended_gcd(a, b)
    (old_r, r) := (a, b)
    (old_s, s) := (1, 0)
    (old_t, t) := (0, 1)

    while r ≠ 0 do
        quotient := old_r div r
        (old_r, r) := (r, old_r − quotient × r)
        (old_s, s) := (s, old_s − quotient × s)
        (old_t, t) := (t, old_t − quotient × t)

    output "Bézout coefficients:", (old_s, old_t)
    output "greatest common divisor:", old_r
    output "quotients by the gcd:", (t, s)
*/
func gcd(a, b int) int {
	oldR, r := a, b
	oldS, s := 1, 0
	oldT, t := 0, 1

	for r != 0 {
		quotient := oldR / r
		oldR, r = r, oldR-quotient*r
		oldS, s = s, oldS-quotient*s
		oldT, t = t, oldT-quotient*t
	}

	//output "Bézout coefficients:", (oldS, oldT)
	//output "greatest common divisor:", oldR
	//output "quotients by the gcd:", (t, s)
	return oldR
}
