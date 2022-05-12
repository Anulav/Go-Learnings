package main

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

type scannedArgs []string

func (a *scannedArgs) Reset(){
	*a = (*a)[0:0]
}

func (a *scannedArgs) 

var ErrClosingQuote = errors.New("Missing closing quote")

func main() {

}
func isQuote(r rune) bool {
	return r == '"' || r == '\''
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
