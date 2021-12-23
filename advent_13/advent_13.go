package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const realFilename = "advent_13.test.txt"

func main() {
	inputFile := readAdvent13File(realFilename)
	result := calcAdvent13Result(inputFile)
	log.Printf("Answer: %s", result.answer)
}

type advent13Result struct {
	answer string
}

func newPhonesCalculator(capacity int) *phonesCalculator {
	return &phonesCalculator{make([]string, 0, capacity), make(map[string]string, capacity)}
}

type phonesCalculator struct {
	fixedPhones      []string
	fixedToSrcPhones map[string]string
}

func (calc *phonesCalculator) isValid() bool {
	return len(calc.fixedPhones) == len(calc.fixedToSrcPhones)
}
func (calc *phonesCalculator) addPhone(phone string) {
	fixedPhone := fixPhone(phone)
	calc.fixedPhones = append(calc.fixedPhones, fixedPhone)
	calc.fixedToSrcPhones[fixedPhone] = phone
	if !calc.isValid() {
		log.Fatalf("calc became invalid after adding %q number (fixed: %q)\n", phone, fixedPhone)
	}
}

func fixPhone(phone string) string {
	cleanBeforeBracketPhone := phone
	//openBracketsAmt := strings.Count(phone, openBracket)
	//if openBracketsAmt != 0 {
	//	split := strings.Split(phone, openBracket)
	//	if len(split) != 2 {
	//		log.Fatalf("phone %q has unexpected brackets amount\n", phone)
	//	}
	//	cleanBeforeBracketPhone = split[1]
	//}
	// replace all non-nums with ""
	nonNumRe := regexp.MustCompile("[^0-9]")
	numbersOnlyPhone := nonNumRe.ReplaceAllString(cleanBeforeBracketPhone, "")
	// check only numbers exist
	numOnlyRe := regexp.MustCompile("^[0-9]+$")
	isFixed := numOnlyRe.Match([]byte(numbersOnlyPhone))
	if !isFixed {
		log.Fatalf("phone %q can not be changed numbers-only: %q\n", phone, numbersOnlyPhone)
	}
	numbersOnlyPhoneLen := len(numbersOnlyPhone)
	if numbersOnlyPhoneLen < 10 {
		log.Fatalf("phone %q number is too short after removing punctuation: %q\n", phone, numbersOnlyPhone)
	}
	charsOf10Number := numbersOnlyPhone[numbersOnlyPhoneLen-10:]
	if len(charsOf10Number) != 10 {
		log.Fatalf("unable to shorten number %q to 10 characters: %q", phone, charsOf10Number)
	}
	// return
	return charsOf10Number
}
func (calc *phonesCalculator) addAllPhones(phones []string) {
	for _, p := range phones {
		calc.addPhone(p)
	}
}
func (calc *phonesCalculator) sort() {
	sort.Strings(calc.fixedPhones)
	//log.Println("sorted phones:")
	//for _, f := range calc.fixedPhones {
	//	fmt.Printf("%v %q %q\n", len(f), f, calc.fixedToSrcPhones[f])
	//}
}

// does not sort
func (calc *phonesCalculator) getPhoneByNumber(num int) string {
	phoneIdx := num - 1
	fixedPhone := calc.fixedPhones[phoneIdx]
	phone := calc.fixedToSrcPhones[fixedPhone]
	return phone
}

func calcAdvent13Result(inputFile advent13File) advent13Result {
	calculator := newPhonesCalculator(inputFile.amount)
	calculator.addAllPhones(inputFile.phones)
	calculator.sort()
	result := advent13Result{calculator.getPhoneByNumber(inputFile.ascNum)}
	return result
}

type advent13File struct {
	amount, ascNum int
	phones         []string
}

func readAdvent13File(filename string) advent13File {
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
	result := advent13File{}
	for lineIdx := 0; lineScanner.Scan(); lineIdx++ {
		lineText := lineScanner.Text()
		switch lineIdx {
		case 0:
			wordScanner := bufio.NewScanner(strings.NewReader(lineText))
			wordScanner.Split(bufio.ScanWords)
			wordScanner.Scan()
			result.amount, _ = strconv.Atoi(wordScanner.Text())
			wordScanner.Scan()
			result.ascNum, _ = strconv.Atoi(wordScanner.Text())
			result.phones = make([]string, 0, result.amount)
		default:
			result.phones = append(result.phones, lineText)
		}
	}
	return result
}
