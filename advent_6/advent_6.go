package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	filename := "advent_6.test.txt"
	inputFile := readAdvent6File(filename)
	log.Printf("File %q contents: %s", filename, inputFile.debugInfo())
	result := calcAdvent6Result(&inputFile)
	log.Printf("Answer: %s", result.print())
}

type advent6Result advent6File // types are same
func (a6r *advent6Result) print() string {
	return string(a6r.directions)
}

// encodes directions in following fashion
// L = +1
// R = -1
// D = +i
// U = -i
// in other words, L/R - real, D/U - imag
type directionsAggregator complex128

func (da *directionsAggregator) addDirections(directions []rune) {
	for _, r := range directions {
		da.addDirection(r)
	}
}

func (da *directionsAggregator) addDirection(r rune) {
	switch r {
	case 'L':
		*da += 1 + 0i
	case 'R':
		*da -= 1 + 0i
	case 'D':
		*da += 0 + 1i
	case 'U':
		*da -= 0 + 1i
	default:
		log.Fatal(fmt.Errorf("unknown rune %v", r))
	}
}

func (da *directionsAggregator) getAggregatedDirections() []rune {
	asComplex := complex128(*da)
	lr := real(asComplex)
	du := imag(asComplex)
	alr := math.Abs(lr)
	adu := math.Abs(du)
	var lsString string
	ilr := int(alr)
	if lr < 0 {
		lsString = strings.Repeat("R", ilr)
	} else if lr > 0 {
		lsString = strings.Repeat("L", ilr)
	} else {
		lsString = ""
	}

	var duString string
	idu := int(adu)
	if du < 0 {
		duString = strings.Repeat("U", idu)
	} else if du > 0 {
		duString = strings.Repeat("D", idu)
	} else {
		duString = ""
	}
	result := duString + lsString
	sort.Strings(strings.Split(result, ""))
	return bytes.Runes([]byte(result))
}

func calcAdvent6Result(inputFile *advent6File) advent6Result {
	aggregator := directionsAggregator(0 + 0i)
	aggregator.addDirections(inputFile.directions)
	return advent6Result{aggregator.getAggregatedDirections()}
}

type advent6File struct {
	directions []rune
}

func (af *advent6File) debugInfo() string {
	ls, rs, ds, us := 0, 0, 0, 0
	for _, r := range (*af).directions {
		switch r {
		case 'R':
			rs++
		case 'L':
			ls++
		case 'U':
			us++
		case 'D':
			ds++
		default:
			log.Fatal(fmt.Errorf("unknown rune %v", r))
		}
	}
	return fmt.Sprintf("ls:%d, rs:%d, ds:%d, us:%d", ls, rs, ds, us)
}

func readAdvent6File(filename string) advent6File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open file %q for reading: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("unable to close file %q: %w", filename, err))
		}
	}()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	result := advent6File{}
	result.directions = make([]rune, 0)
	for scanner.Scan() {
		result.directions = append(result.directions, bytes.Runes([]byte(scanner.Text()))[0])
	}

	return result
}
