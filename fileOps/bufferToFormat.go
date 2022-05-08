package main

import (
	"bytes"
	"fmt"
	"os"
)

type book struct {
	Author, Title string
	Year          int
}

func main() {
	const grr string = "G.R.R"
	//	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	//	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.

	dst, err := os.OpenFile("test_bTF.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("\n\nError", err)
		return
	}
	defer dst.Close()

	bookList := []book{
		{Author: grr, Title: "A Game of Thrones", Year: 1996},
		{Author: grr, Title: "A Clash of Kings", Year: 1998},
		{Author: grr, Title: "A Storm of Swords", Year: 2000},
		{Author: grr, Title: "A Feast for Crows", Year: 2005},
		{Author: grr, Title: "A Dance with Dragons", Year: 2011},
		// if year is omitted it defaulting to zero value
		{Author: grr, Title: "The Winds of Winter"},
		{Author: grr, Title: "A Dream of Spring"},
	}

	b := bytes.NewBuffer(make([]byte, 0, 16))
	for _, v := range bookList {
		fmt.Fprintf(b, "%s - %s", v.Title, v.Author)

		if v.Year > 0 {
			fmt.Fprintf(b, " (%d)", v.Year)
		}
		b.WriteRune('\n')
		if _, err := b.WriteTo(dst); err != nil {
			fmt.Println("\n\nError", err)
			return
		}
	}

}
