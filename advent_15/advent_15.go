package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
)

const realFilename = "advent_15.test.txt"

func main() {
	inputFile := readAdvent15File(realFilename)
	result := calcAdvent15Result(inputFile)
	log.Printf("Answer: %d\n", result.answer)
}

type advent15Result struct {
	answer int
}

func calcAdvent15Result(inputFile advent15File) advent15Result {
	mmAgg := newMinMaxAggregator()
	for _, p := range inputFile.points {
		mmAgg.add(p)
	}

	minXAgg := newMinPerAxisAggregator()
	for x := mmAgg.minX; x <= mmAgg.maxX; x++ {
		p := point{x, 0}
		d := p.sumXDistanceTo(inputFile.points...)
		minXAgg.add(x, d)
	}
	minYAgg := newMinPerAxisAggregator()
	for y := mmAgg.minY; y <= mmAgg.maxY; y++ {
		p := point{0, y}
		d := p.sumYDistanceTo(inputFile.points...)
		minYAgg.add(y, d)
	}
	sepPoint := point{minXAgg.minValue(), minYAgg.minValue()}
	result := advent15Result{sepPoint.abs()}
	return result
}

func newMinMaxAggregator() minMaxAggregator {
	return minMaxAggregator{math.MaxInt, math.MinInt, math.MaxInt, math.MinInt}
}

type minMaxAggregator struct {
	minX, maxX, minY, maxY int
}

func (mma *minMaxAggregator) add(p point) {
	mma.minX = minInt(mma.minX, p.x)
	mma.maxX = maxInt(mma.maxX, p.x)
	mma.minY = minInt(mma.minY, p.y)
	mma.maxY = maxInt(mma.maxY, p.y)
}

func sum(x ...uint64) uint64 {
	r := uint64(0)
	for _, s := range x {
		r = sum0(r, s)
	}
	return r
}
func sum0(a, b uint64) uint64 {
	if res, carry := bits.Add64(a, b, 0); carry != 0 {
		log.Fatalln(fmt.Errorf("addition overflow"))
	} else {
		return res
	}
	//goland:noinspection GoUnreachableCode
	panic("should've never get here")
}

func xDistance(p1, p2 point) uint64 {
	return intAbs(p1.x - p2.x)
}
func yDistance(p1 point, p2 point) uint64 {
	return intAbs(p1.y - p2.y)
}

type point struct {
	x, y int
}

func (p point) xDistanceTo(a point) uint64 {
	return xDistance(p, a)
}
func (p point) yDistanceTo(a point) uint64 {
	return yDistance(p, a)
}
func (p point) sumXDistanceTo(ps ...point) uint64 {
	result := uint64(0)
	for _, a := range ps {
		result = sum(result, p.xDistanceTo(a))
	}
	return result
}
func (p point) sumYDistanceTo(ps ...point) uint64 {
	result := uint64(0)
	for _, a := range ps {
		result = sum(result, p.yDistanceTo(a))
	}
	return result
}
func (p point) sumDistanceTo(ps ...point) uint64 {
	return sum(p.sumXDistanceTo(ps...), p.sumYDistanceTo(ps...))
}
func (p point) abs() int {
	return p.y + p.x
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func intAbs(i int) uint64 {
	if i < 0 {
		return uint64(-i)
	}
	return uint64(i)
}

func newMinPerAxisAggregator() minPerAxisAggregator {
	return minPerAxisAggregator{minDistance: math.MaxUint64}
}

type minPerAxisAggregator struct {
	minDistance      uint64
	axisValuesForMin []int
}

func (aa *minPerAxisAggregator) add(axisValue int, d uint64) {
	if d < aa.minDistance {
		aa.minDistance = d
		aa.axisValuesForMin = []int{axisValue}
	} else if d == aa.minDistance {
		aa.axisValuesForMin = append(aa.axisValuesForMin, axisValue)
	}
}
func (aa minPerAxisAggregator) minValue() int {
	values := aa.axisValuesForMin
	sort.Ints(values)
	return values[0]
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
				p.y = number
				result.points = append(result.points, p)
			}
		}
	}

	return result
}

type advent15File struct {
	pointsAmount int
	points       []point
}
