package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
)

const realFilename = "advent_15.test.txt"

func main() {
	inputFile := readAdvent15File(realFilename)
	result := calcAdvent15Result(inputFile)
	log.Printf("Answer: %d\n", result.answer)
}

func readAdvent15File(filename string) advent15File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(fmt.Errorf("error opening file %q: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("error closing file %q: %w", filename, err))
		}
	}()

	wordScanner := bufio.NewScanner(file)
	wordScanner.Split(bufio.ScanWords)
	result := advent15File{}

	var p point
	for wordIdx := 0; wordScanner.Scan(); wordIdx++ {
		wordText := wordScanner.Text()
		number, err := strconv.Atoi(wordText)
		if err != nil {
			lineNumber := wordIdx/2 + 1
			log.Fatalln(fmt.Errorf("nan %s at line %d: %w", wordText, lineNumber, err))
		}
		switch wordIdx {
		case 0:
			result.pointsAmount = number
			result.points = make([]point, 0, number)
		default:
			if wordIdx%2 == 1 {
				p = point{number, 0}
			} else {
				p.i = number
				result.points = append(result.points, p)
			}
		}
	}

	return result
}

func intAbs(i int) uint64 {
	if i < 0 {
		return uint64(-i)
	}
	return uint64(i)
}
func distance(p1, p2 point) uint64 {
	return sum(intAbs(p1.i-p2.i), intAbs(p1.r-p2.r))
}

type point struct {
	r /*x*/, i /*y*/ int
}

func (p point) distanceTo(a point) uint64 {
	return distance(p, a)
}
func (p point) sumDistance(ps ...point) uint64 {
	result := uint64(0)
	for _, a := range ps {
		result = sum(result, p.distanceTo(a))
	}
	return result
}

func (p *point) asComplex() complex128 {
	return complex(float64(p.r), float64(p.i))
}

type advent15File struct {
	pointsAmount int
	points       []point
}

type advent15Result struct {
	answer int
}

func calcAdvent15Result(inputFile advent15File) advent15Result {
	result := advent15Result{}
	return result
}

func sum(a, b uint64) uint64 {
	if res, carry := bits.Add64(a, b, 0); carry != 0 {
		log.Fatalln(fmt.Errorf("addition overflow"))
	} else {
		return res
	}
	//goland:noinspection GoUnreachableCode
	panic("should've never get here")
}
