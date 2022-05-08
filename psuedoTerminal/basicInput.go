package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	fmt.Fprint(w, "Write something\n")
	for {
		s.Scan() //Get the next token
		msg := string(s.Bytes())
		if msg == "exit" {
			return
		}
		fmt.Fprint(w, "You wrote ")
		w.WriteString(msg)
		fmt.Fprint(w, "\n")

	}
}
