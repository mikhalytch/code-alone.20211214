package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const realFilename = "advent_14.test.txt"

func main() {
	inputFile := readAdvent14File(realFilename)
	result := calcAdvent14Result(inputFile)
	log.Printf("Result: %q\n", result.answer)
}

type advent14Result struct {
	answer string
}

func calcAdvent14Result(inputFile advent14File) advent14Result {
	rules := inputFile.rules
	perPosRestrictions := groupPositionRules(rules)

	result := make([]rune, 0, inputFile.codeLength)
	for codeNum := 1; codeNum <= inputFile.codeLength; codeNum++ {
		result = append(result, calcMedian(createLimitedSet(perPosRestrictions[codeNum])))
	}

	return advent14Result{string(result)}
}

func groupPositionRules(rules []simpleRestriction) map[int][]rune {
	perPosRestrictions := make(map[int][]rune, len(rules))
	for _, r := range rules {
		position := r.codePosition
		if _, ok := perPosRestrictions[position]; !ok {
			perPosRestrictions[position] = make([]rune, 0)
		}
		perPosRestrictions[position] = append(perPosRestrictions[position], r.restrictedSymbol)
	}
	return perPosRestrictions
}

func calcMedian(set []rune) rune {
	stringSet := make([]string, 0, len(set))
	for _, s := range set {
		stringSet = append(stringSet, string(s))
	}
	sort.Strings(stringSet)
	stringsAmt := float64(len(stringSet))
	for stringIdx, str := range stringSet {
		stringsNumber := float64(stringIdx + 1)
		if stringsNumber/stringsAmt >= 0.5 {
			return bytes.Runes([]byte(str))[0]
		}
	}
	return 0
}

var fullSetM = map[rune]bool{'A': true, 'B': true, 'C': true, 'D': true, 'E': true, 'F': true, 'G': true, 'H': true,
	'I': true, 'J': true, 'K': true, 'L': true, 'M': true, 'N': true, 'O': true, 'P': true, 'Q': true, 'R': true,
	'S': true, 'T': true, 'U': true, 'V': true, 'W': true, 'X': true, 'Y': true, 'Z': true}
var fullSet = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func createLimitedSet(restricted []rune) []rune {

	result := make([]rune, 0)

	restrictions := make(map[rune]bool, len(restricted))
	for _, r := range restricted {
		restrictions[r] = true

		if !fullSetM[r] { // check if restricted was present at full set
			log.Fatalln(fmt.Errorf("restriction %q wasn't present at full set", r)) // error otherwise
		}
	}

	for _, f := range fullSet {
		if !restrictions[f] {
			result = append(result, f)
		}
	}

	return result
}

type simpleRestriction struct {
	codePosition     int
	restrictedSymbol rune
}

type advent14File struct {
	codeLength, rulesAmount int
	rules                   []simpleRestriction
}

func readAdvent14File(filename string) advent14File {
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

	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	result := advent14File{}

	for lineIdx := 0; lineScanner.Scan(); lineIdx++ {
		lineText := lineScanner.Text()
		wordScanner := bufio.NewScanner(strings.NewReader(lineText))
		wordScanner.Split(bufio.ScanWords)

		wordScanner.Scan()
		fWord := wordScanner.Text()
		fNum, err := strconv.Atoi(fWord)
		if err != nil {
			log.Fatalln(fmt.Errorf("nan %q lineIdx %d word 1: %w", fWord, lineIdx, err))
		}
		wordScanner.Scan()
		sWord := wordScanner.Text()

		switch lineIdx {
		case 0:
			result.codeLength = fNum
			result.rulesAmount, err = strconv.Atoi(sWord)
			if err != nil {
				log.Fatalln(fmt.Errorf("nan %q lineIdx %d word 2: %w", sWord, lineIdx, err))
			}
			result.rules = make([]simpleRestriction, 0, result.rulesAmount)
		default:
			if len(sWord) > 1 {
				log.Fatalln(fmt.Errorf("not a rune (line %d): %s", lineIdx, sWord))
			}
			restriction := simpleRestriction{fNum, bytes.Runes([]byte(sWord))[0]}
			result.rules = append(result.rules, restriction)
		}
	}

	return result
}
