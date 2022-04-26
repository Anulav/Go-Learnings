package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give in an file name")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f) //wrapping reader with a buffered one.

	var rowCount int
	for err == nil {
		var b []byte
		for moar := true; err == nil && moar; {
			b, moar, err = r.ReadLine()
			if err == nil {
				fmt.Println(string(b))
				rowCount++
			}
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println("\n\nError", err)
		return
	}
	fmt.Println("\n rowCounts ", rowCount)
}
