package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

type cmdstr struct {
	Name, Description string
	Action            func(w io.Writer, args []string) bool
}

func (c cmdstr) match(cmdName string) bool {
	return cmdName == c.Name
}

func (c cmdstr) run(w io.Writer, args []string) bool {
	return c.Action(w, args)
}

func main() {
	w := os.Stdout
	r := bufio.NewScanner(os.Stdin)
	fmt.Fprintln(w, "Input a command")
	r.Scan()
	args := strings.Split(string(r.Bytes()), " ")
	cmd := args[0]
	args = args[1:]
	var cmds []cmdstr
	help := cmdstr{
		Name:        "help",
		Description: "Shows all available commands with short description",
		Action: func(w io.Writer, args []string) bool {
			for _, c := range cmds {
				fmt.Fprintf(w, " - %-15s %s\n", c.Name, c.Description)
			}
			return false
		},
	}

	shuffle := cmdstr{
		Name:        "shuffle",
		Description: "Shuffles args",
		Action:      shuffleCmd,
	}

	greeting := cmdstr{
		Name:        "greeting",
		Description: "Greets passed name of the user",
		Action:      greetingCmd,
	}

	printFile := cmdstr{
		Name:        "printf",
		Description: "Prints file passed in arg",
		Action:      printFileCmd,
	}

	idx := -1
	cmds = append(cmds, help, shuffle, printFile, greeting)
	for i := range cmds {
		if cmds[i].match(cmd) {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Fprintln(w, "Command not found. Use \"help\", to know more about available commands")
		return
	}
	if cmds[idx].run(w, args) {
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
