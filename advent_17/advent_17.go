package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

const realFilename = "advent_17.test.txt"

func main() {
	inputFile := readAdvent17File(realFilename)
	result := calcAdvent17Result(inputFile)
	log.Printf("Answer: %d", result.answer)
}

type advent17Result struct {
	answer int
}

type zeroLevelBracesCalculator struct {
	amountClosed  int
	currentlyOpen bool
}

func (calc *zeroLevelBracesCalculator) onOpen(to *realParenStack) {
	if to.level() == 1 && to.p == '{' {
		calc.currentlyOpen = true
	}
}
func (calc *zeroLevelBracesCalculator) onClose(from *realParenStack, p paren) {
	if calc.currentlyOpen && p == '}' && from.level() == 1 {
		calc.currentlyOpen = false
		calc.amountClosed += 1
	}
}
func (calc *zeroLevelBracesCalculator) onError() {
	calc.currentlyOpen = false
}

type paren rune

func unmatchedClosingParenError(p paren) error {
	return fmt.Errorf("unmatched closing paren %s", string(p))
}

// returns error in case of unmatched closing paren
func (p paren) isClosing() bool {
	if p == '(' || p == '[' || p == '{' {
		return false
	}
	return true
}
func (p paren) isComplemetary(o paren) (bool, error) {
	if p.isClosing() {
		return false, unmatchedClosingParenError(o)
	}
	switch p {
	case '(':
		return o == ')', nil
	case '[':
		return o == ']', nil
	case '{':
		return o == '}', nil
	}
	panic("shouldn't get here")
}

type parenStack interface {
	add(p paren, calc *zeroLevelBracesCalculator) parenStack
	level() int
	runesTo(buf *[]rune)
}

func createEmptyParenStack() parenStack {
	return &emptyParenStack{}
}

type emptyParenStack struct{}

func (e *emptyParenStack) level() int {
	return 0
}

func (e *emptyParenStack) runesTo(_ *[]rune) {
	// nop
}

func (e *emptyParenStack) add(p paren, calc *zeroLevelBracesCalculator) parenStack {
	if p.isClosing() {
		return createEmptyParenStack()
	}
	result := &realParenStack{e, p, 1}
	calc.onOpen(result)
	return result
}

func (e *emptyParenStack) String() string {
	return fmt.Sprintf("EMPTY_PAREN_STACK")
}

type realParenStack struct {
	prev parenStack
	p    paren
	l    int
}

func (ps *realParenStack) add(p paren, calc *zeroLevelBracesCalculator) parenStack {
	complemetary, err := ps.p.isComplemetary(p)
	if err != nil {
		calc.onError()
		return createEmptyParenStack()
	}
	if complemetary {
		calc.onClose(ps, p)
		return ps.prev
	} else {
		result := &realParenStack{ps, p, ps.l + 1}
		calc.onOpen(result)
		return result
	}
}

func (ps *realParenStack) level() int {
	return ps.l
}

func (ps *realParenStack) runesTo(buf *[]rune) {
	ps.prev.runesTo(buf) // recurse
	*buf = append(*buf, rune(ps.p))
}
func (ps *realParenStack) String() string {
	var runes []rune
	ps.runesTo(&runes)
	return fmt.Sprintf("%q", string(runes))
}

func calcAdvent17Result(inputFile advent17File) advent17Result {
	calc := &zeroLevelBracesCalculator{}
	stack := createEmptyParenStack()
	for _, r := range inputFile.symbols {
		p := paren(r)
		stack = stack.add(p, calc)
	}
	return advent17Result{calc.amountClosed}
}

type advent17File struct {
	symbols []rune
}

func readAdvent17File(filename string) advent17File {
	str, err := os.ReadFile(filename)
	if err != nil {
		log.Println(fmt.Errorf("error opening file %q: %w", filename, err))
	}
	return advent17File{bytes.Runes([]byte(strings.TrimSpace(string(str))))}
}
