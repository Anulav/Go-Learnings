package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

var cmdExec func(w io.Writer, args []string) (exit bool)

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
		cmdExec = exitCmd
	case "hello":
		cmdExec = greetingCmd
	case "shuffle":
		cmdExec = shuffleCmd
	case "printf":
		cmdExec = printFileCmd
	default:
	}
	if cmdExec == nil {
		fmt.Fprintf(w, "The cmd is not valid. Exiting")
	}
	if cmdExec(w, args) { // execute and exit if true
		return
	}
}
func greetingCmd(w io.Writer, args []string) bool {
	fmt.Fprintf(w, "\nHello, %s\n", args[0])
	return true
}

func exitCmd(w io.Writer, args []string) bool {
	fmt.Fprintf(w, "\nBye. Have a nice day\n")
	return true
}

func shuffleCmd(w io.Writer, args []string) bool {
	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
	})
	for i := range args {
		if i > 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprintf(w, "%s", args[i])
	}
	fmt.Fprintln(w)
	return false
}

func printFileCmd(w io.Writer, args []string) bool {
	if len(args) != 1 {
		fmt.Print(w, "\nPlease give at least one file name")
		return false
	}
	r, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(w, "\nCan't open the file %s due to: %s", args[0], err)
	}
	defer r.Close()
	if _, err := io.Copy(w, r); err != nil {
		fmt.Fprintf(w, "\nCan't print the contents of the file %s due to: %s ", args[0], err)
	}
	fmt.Fprintln(w)
	return false
}
