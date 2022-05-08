package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Enter a valid file name and some content")
		return
	}
	if err := ioutil.WriteFile(os.Args[1], []byte(os.Args[2]), 0644); err != nil {
		fmt.Println("\n\nError : ", err)
	}
}
