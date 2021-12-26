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

	log.Println("branches traversed", total)

	return advent16Result{minPathsAgg.curLen}
}

type pathCalculationStats struct {
	finished, tooLongAlready, tooLongAtXRoad, loopDetected, eyeletDetected, noMoreMoves uint64
}

func (p pathCalculationStats) total() uint64 {
	return p.finished + p.tooLongAlready + p.tooLongAtXRoad + p.loopDetected + p.eyeletDetected + p.noMoreMoves
}
func (p pathCalculationStats) String() string {
	return fmt.Sprintf(
		"total:%d {finished:%d, tooLongAlready:%d, tooLongAtXRoad:%d, loopDetected:%d, eyeletDetected:%d, noMoreMoves:%d}",
		p.total(), p.finished, p.tooLongAlready, p.tooLongAtXRoad, p.loopDetected, p.eyeletDetected, p.noMoreMoves,
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
	var backPoint *point = nil
	if path.nose != nil {
		backPoint = &path.nose.tail
	}
	moves := f.getPossibleMoves(path.tail, backPoint)
	if len(moves) > 1 { // cross-road case
		if path.hasAnyOfPoints(moves...) {
			total.eyeletDetected++
			return
		}

		currentXRoadAgg, ok := xRoadsPathAgg[path.tail]
		if !ok {
			currentXRoadAgg = *newMinAggregator()
		}
		if currentXRoadAgg.curLen <= currentPathLength { // we've been here already, with shorter path
			total.tooLongAtXRoad++
			return
		} else {
			currentXRoadAgg.addPath(*path)
		}
		xRoadsPathAgg[path.tail] = currentXRoadAgg
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
	rows          []fieldRow
	start, finish point
	stats         fieldStats
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
func (f *field) getOrderedPointPermutations(p point) []point {
	addYBasedOnYDiff := func(yDiff int) []point {
		if yDiff < 0 {
			return []point{{p.x, p.y + 1}, {p.x, p.y - 1}}
		} else {
			return []point{{p.x, p.y - 1}, {p.x, p.y + 1}}
		}
	}
	addXBasedOnXDiff := func(xDiff int) []point {
		if xDiff < 0 {
			return []point{{p.x + 1, p.y}, {p.x - 1, p.y}}
		} else {
			return []point{{p.x - 1, p.y}, {p.x + 1, p.y}}
		}
	}
	result := make([]point, 0)
	//result := []point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}}
	xDiff := p.x - f.finish.x
	yDiff := p.y - f.finish.y
	if intAbs(xDiff) > intAbs(yDiff) {
		if xDiff < 0 {
			result = append(result, point{p.x + 1, p.y})
			result = append(result, addYBasedOnYDiff(yDiff)...)
			result = append(result, point{p.x - 1, p.y})
		} else {
			result = append(result, point{p.x - 1, p.y})
			result = append(result, addYBasedOnYDiff(yDiff)...)
			result = append(result, point{p.x + 1, p.y})

		}
	} else {
		if yDiff < 0 {
			result = append(result, point{p.x, p.y + 1})
			result = append(result, addXBasedOnXDiff(xDiff)...)
			result = append(result, point{p.x, p.y - 1})
		} else {
			result = append(result, point{p.x, p.y - 1})
			result = append(result, addXBasedOnXDiff(xDiff)...)
			result = append(result, point{p.x, p.y + 1})
		}
	}
	return result
}
func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func (f *field) getPossibleMoves(p point, backPoint *point) []point {
	filter := func(perm point) bool {
		return (backPoint != nil && perm != *backPoint) || (backPoint == nil)
	}
	var result []point
	for _, perm := range f.getOrderedPointPermutations(p) {
		if filter(perm) && f.isWalkable(perm) {
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
			case finish:
				result.finish = point{colIdx, rowIdx}
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

func (pp *path) hasThePoint(p *point) bool {
	for current := pp; current != nil; current = current.nose {
		if current.tail == *p {
			return true
		}
	}
	return false
}
func (pp *path) hasAnyOfPoints(points ...point) bool {
	for _, p := range points {
		if pp.hasThePoint(&p) {
			return true
		}
	}
	return false
}

// returns false in case of loop
func (pp path) addPoint(p point) (*path, bool) {
	newPath := path{&pp, p, pp.len + 1}
	loopIdentified := pp.hasThePoint(&p)
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

type point struct {
	// coordinates; left upper corner is 0,0
	// x - col
	// y - row
	x, y int
}
