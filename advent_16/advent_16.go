package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const realFilename = "advent_16.test.txt"

func main() {

}

type fieldPosition rune

const (
	road  fieldPosition = '.'
	wall  fieldPosition = '#'
	start fieldPosition = 'A'
	end   fieldPosition = 'B'
)

type advent16File struct {
	f field
}
type field struct {
	rows  []fieldRow
	start point
}
type fieldRow struct {
	positions []fieldPosition
}

func readAdvent16File(filename string) advent16File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to open file %q: %w", filename, err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(fmt.Errorf("error closing file %q: %w", filename, err))
		}
	}()

	rowScanner := bufio.NewScanner(file)
	rowScanner.Split(bufio.ScanLines)
	result := field{}
	for rowIdx := 0; rowScanner.Scan(); rowIdx++ {
		rowText := rowScanner.Text()
		colScanner := bufio.NewScanner(strings.NewReader(rowText))
		colScanner.Split(bufio.ScanWords)
		row := fieldRow{}
		for colIdx := 0; colScanner.Scan(); colIdx++ {
			colText := colScanner.Text()
			col := fieldPosition(bytes.Runes([]byte(colText))[0])
			row.positions = append(row.positions, col)
			switch col {
			case start:
				result.start = point{colIdx, rowIdx}
			}
		}
		result.rows = append(result.rows, row)
	}

	return advent16File{result}
}

// aggregator
func newMinPathsAggregator() minPathsAggregator {
	return minPathsAggregator{curLen: math.MaxInt}
}

type minPathsAggregator struct {
	paths  []path
	curLen int
}

func (pa *minPathsAggregator) addPath(p path) {
	l := p.len()
	if l < pa.curLen {
		pa.paths = []path{p}
		pa.curLen = l
	} else if l == pa.curLen {
		pa.paths = append(pa.paths, p)
	}
}

type path struct {
	points         []point
	pointsRegistry map[point]bool
}

// returns false in case of loop
func (pp *path) addPoint(p point) bool {
	pp.points = append(pp.points, p)
	if pp.pointsRegistry[p] {
		return false
	}
	pp.pointsRegistry[p] = true
	return true
}
func (pp path) len() int {
	return len(pp.points)
}

type point struct {
	// coordinates; left upper corner is 0,0
	// x - col
	// y - row
	x, y int
}
