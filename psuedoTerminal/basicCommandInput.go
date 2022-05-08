package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := os.Stdout
	r := bufio.NewScanner(os.Stdin)
	fmt.Fprintln(w, "Input a command")
	r.Scan()
	args := strings.Split(string(r.Bytes()), " ")
	cmd := args[0]
	args = args[1:]

	switch cmd { //Cases do not fall through from one to the next as in C-like languages (though there is a rarely used fallthrough statement that overrides this behavior).
	case "exit":
		return
	case "hello":
		fmt.Fprintf(w, "Hi there with option %s", args)
	default:
		fmt.Fprintf(w, "Your entered command %s is not valid. Exiting \n", cmd)

	}
}
