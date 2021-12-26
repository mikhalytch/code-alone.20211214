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
	xRoadsPathsAgg := make(map[point]minAggregator, 0)
	initialPath := createPathStart(inputFile.f.start)
	stepRecursively(&inputFile.f, minPathsAgg, xRoadsPathsAgg, initialPath)

	log.Println("Path", minPathsAgg.paths[0].asSliceOfPointsReversed())
	log.Println("branches traversed", total)

	return advent16Result{minPathsAgg.curLen}
}

type pathCalculationStats struct {
	finished, tooLongAlready, tooLongAtXRoad, loopDetected, noMoreMoves uint64
}

func (p pathCalculationStats) total() uint64 {
	return p.finished + p.tooLongAlready + p.tooLongAtXRoad + p.loopDetected + p.noMoreMoves
}
func (p pathCalculationStats) String() string {
	return fmt.Sprintf(
		"total:%d {finished:%d, tooLongAlready:%d, tooLongAtXRoad:%d, loopDetected:%d, noMoreMoves:%d}",
		p.total(), p.finished, p.tooLongAlready, p.tooLongAtXRoad, p.loopDetected, p.noMoreMoves,
	)
}

var total = pathCalculationStats{}

func stepRecursively(f *field, minPathsAgg *minPathsAggregator, xRoadsPathAgg map[point]minAggregator, path *path) {
	currentPathLength := path.length()
	if currentPathLength > f.stats.movablePositions {
		log.Fatalln("path is too long, should've looped long ago")
		//return // too long (however strange this has not looped)
	}
	if currentPathLength > minPathsAgg.curLen {
		total.tooLongAlready++
		return // already too long
	}
	if f.isFinish(path.tail) {
		minPathsAgg.addPath(*path)
		total.finished++
		return
	}
	filter := make(map[point]bool, 0)
	if path.nose != nil {
		filter[path.nose.tail] = true
	}
	moves := f.getPossibleMoves(path.tail, filter)
	if len(moves) > 1 { // cross-road case
		rowNum := path.tail.y + 1
		colNum := path.tail.x + 1
		if _, ok := xRoadsPathAgg[path.tail]; !ok {
			xRoadsPathAgg[path.tail] = *newMinAggregator()
		}
		agg := xRoadsPathAgg[path.tail]
		if agg.curLen <= currentPathLength { // we've been here already, with shorter path
			if rowNum > 95 && colNum < 5 {
				log.Printf("Re-Entered crossroad at rowNum:%d colNum:%d (len: %d; answers:%d[%d]); moves: %v \n",
					rowNum, colNum, path.len, len(minPathsAgg.paths), minPathsAgg.curLen, moves)
			}
			total.tooLongAtXRoad++
			return
		} else {
			if rowNum > 95 && colNum < 5 {
				log.Printf("crossroad at rowNum:%d colNum:%d (len: %d; answers:%d[%d]); moves: %v \n",
					rowNum, colNum, path.len, len(minPathsAgg.paths), minPathsAgg.curLen, moves)
			}
			agg.addPath(*path)
		}
		xRoadsPathAgg[path.tail] = agg
	}
	for _, m := range moves {
		if newPath, ok := path.addPoint(m); ok {
			stepRecursively(f, minPathsAgg, xRoadsPathAgg, newPath)
		} else {
			// otherwise - skip this branch, since loop detected
			total.loopDetected++
		}
	}
	total.noMoreMoves++ // no moves case
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
	stats fieldStats
}
type fieldStats struct {
	rowNum, colNum, size, movablePositions int
}
type fieldRow struct {
	positions []fieldPosition // columns
}

func (f *field) recalculateStats() {
	positions := 0
	rows := 0
	m := 0
	for rowIdx, row := range f.rows {
		rows++
		for _, pos := range row.positions {
			if rowIdx == 0 {
				positions++
			}
			if f.isPositionWalkable(pos) {
				m++
			}
		}
	}
	f.stats = fieldStats{rows, positions, rows * positions, m}
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
			position := f.positionAt(p)
			if f.isPositionWalkable(position) {
				return true
			}
		}
	}
	return false
}

func (f *field) isPositionWalkable(position fieldPosition) bool {
	if position != wall {
		return true
	}
	return false
}
func (f *field) getPossibleMoves(p point, filter map[point]bool) []point {
	var result []point
	permutations := []point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}}
	for _, perm := range permutations {
		if !filter[perm] && f.isWalkable(perm) {
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
	result.recalculateStats()
	return advent16File{result}
}

// aggregator
func newMinAggregator() *minAggregator {
	return &minAggregator{curLen: math.MaxInt}
}

type minAggregator struct {
	curLen int
}

func (pa *minAggregator) addPath(p path) {
	l := p.length()
	if l < pa.curLen {
		pa.curLen = l
	}
}
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

func (pp path) asSliceOfPointsReversed() []point {
	var result []point
	for current := &pp; current != nil; current = current.nose {
		result = append(result, current.tail)
	}
	return result
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
