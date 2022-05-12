package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrClosingQuote = errors.New("Missing closing quote")

func isQuote(r rune) bool {
	return r == '"' || r == '\''
}
func main() {
	s := bufio.NewScanner(strings.NewReader(`And the 'cats in 
	the craddle' and "the silver's spoon" nasa put a men on "the moon`))
	s.Split(ScanArgs)
	for s.Scan() {
		fmt.Printf("%q\n", s.Text())
	}
	fmt.Println(s.Err())
}

func ScanArgs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start, first := 0, rune(0)
	for width := 0; start < len(data); start += width {
		first, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(first) {
			break
		}
	}
	if isQuote(first) {
		start++
	}
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if ok := isQuote(first); !ok && unicode.IsSpace(r) || ok && r == first {
			return i + width, data[start:i], nil
		}
	}
	if atEOF && len(data) > start {
		if isQuote(first) {
			err = ErrClosingQuote
		}
		return len(data), data[start:], err
	}
	if isQuote(first) {
		start--
	}
	return start, nil, nil
}
