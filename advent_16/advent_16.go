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
	inputFile := readAdvent16File(realFilename)
	result := calcAdvent16Result(inputFile)
	log.Printf("Answer: %d\n", result.answer)
}

type advent16Result struct {
	answer int
}

func calcAdvent16Result(inputFile advent16File) advent16Result {
	minPathsAgg := newMinPathsAggregator()
	initialPath := createPathStart(inputFile.f.start)
	stepRecursively(&inputFile.f, minPathsAgg, initialPath)

	return advent16Result{minPathsAgg.curLen}
}

func stepRecursively(f *field, minPathsAgg *minPathsAggregator, path *path) {
	pathLength := path.length()
	if pathLength > f.size() {
		log.Fatalln("path is too long, should've looped long ago")
		//return // too long (however strange this has not looped)
	}
	if pathLength > minPathsAgg.curLen {
		return // already too long
	}
	if f.isFinish(path.tail) {
		minPathsAgg.addPath(*path)
		return
	}
	moves := f.getPossibleMoves(path.tail)
	for _, m := range moves {
		if newPath, ok := path.addPoint(m); ok {
			stepRecursively(f, minPathsAgg, newPath)
		} // otherwise - skip, since loop detected
	}
}

type fieldPosition rune

const (
	_/*road*/ fieldPosition = '.'
	wall                    fieldPosition = '#'
	start                   fieldPosition = 'A'
	finish                  fieldPosition = 'B'
)

type advent16File struct {
	f field
}
type field struct {
	rows  []fieldRow
	start point
}
type fieldRow struct {
	positions []fieldPosition // columns
}

func (f *field) size() int {
	return len(f.rows[0].positions) * len(f.rows)
}
func (f *field) positionAt(p point) fieldPosition {
	return f.rows[p.y].positions[p.x]
}
func (f *field) isFinish(p point) bool {
	return f.positionAt(p) == finish
}
func (f *field) isWalkable(p point) bool {
	if p.x >= 0 && p.x < len(f.rows[0].positions) {
		if p.y >= 0 && p.y < len(f.rows) {
			if f.positionAt(p) != wall {
				return true
			}
		}
	}
	return false
}
func (f *field) getPossibleMoves(p point) []point {
	var result []point
	permutations := []point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}}
	for _, perm := range permutations {
		if f.isWalkable(perm) {
			result = append(result, perm)
		}
	}
	return result
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
func newMinPathsAggregator() *minPathsAggregator {
	return &minPathsAggregator{curLen: math.MaxInt}
}

type minPathsAggregator struct {
	paths  []path
	curLen int
}

func (pa *minPathsAggregator) addPath(p path) {
	l := p.length()
	if l < pa.curLen {
		pa.paths = []path{p}
		pa.curLen = l
	} else if l == pa.curLen {
		pa.paths = append(pa.paths, p)
	}
}

func createPathStart(start point) *path {
	return &path{nil, start, 0}
}

// tail
type path struct {
	nose *path
	tail point
	len  int // nose.len + 1
}

func (pp path) pointsRegistry() map[point]bool {
	result := make(map[point]bool, pp.len)
	for current := &pp; current != nil; current = current.nose {
		result[current.tail] = true
	}
	return result
}

// returns false in case of loop
func (pp path) addPoint(p point) (*path, bool) {
	newPath := path{&pp, p, pp.len + 1}
	loopIdentified := pp.pointsRegistry()[p]
	return &newPath, !loopIdentified
}

// minus starting node
func (pp path) length() int {
	return pp.len
}

/*type path struct {
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
}*/

type point struct {
	// coordinates; left upper corner is 0,0
	// x - col
	// y - row
	x, y int
}
