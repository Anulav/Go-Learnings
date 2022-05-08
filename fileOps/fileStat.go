package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter a valid filename")
		return
	}
	f, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Print("\nError: ", err)
		return
	}
	fmt.Println(f.ModTime())

}
