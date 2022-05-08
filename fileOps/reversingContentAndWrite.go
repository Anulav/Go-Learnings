package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Enter a valid filename")
		return
	}
	file, err := os.Open(os.Args[1]) //Opening file
	if err != nil {
		fmt.Println("\n\nError", err)
		return
	}

	defer file.Close()

	dst, err := os.Open(os.Args[2]) //Opening file
	if err != nil {
		fmt.Println("\n\nError", err)
		return
	}

	defer dst.Close()

	cur, err := file.Seek(0, os.SEEK_END) //Seeking till end
	if err != nil {
		fmt.Println("\n\nError", err)
		return
	}
	fmt.Println("0. cur is ", cur)
	b := make([]byte, 16) //Making slica for a buffer

	for step, r, w := int64(16), 0, 0; cur != 0; {
		fmt.Println("cur is", cur)
		fmt.Println("step is", step)
		if cur < step {
			b, step = b[:cur], cur
		}
		cur = cur - step
		fmt.Println("2. cur is", cur)
		fmt.Println("2. step is", step)

		_, err = file.Seek(0, os.SEEK_SET)
		if err != nil {
			fmt.Println("\n\nError", err)
			break
		}
		if r, err = file.Read(b); err != nil || r != len(b) {
			if err == nil {
				err = fmt.Errorf("read expected %d bytes, got %d", len(b), r)
			}
			break
		}
		fmt.Println("r is: ", r)
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			switch { // Swap (\r\n) so they get back in place
			case b[i] == '\r' && b[i+1] == '\n':
				b[i], b[i+1] = b[i+1], b[i]
			case j != len(b)-1 && b[j-1] == '\r' && b[j] == '\n':
				b[j], b[j-1] = b[j-1], b[j]
			}
			b[i], b[j] = b[j], b[i] // swap bytes
		}
		if w, err = dst.Write(b); err != nil || w != len(b) {
			if err != nil {
				err = fmt.Errorf("write: expected %d bytes, got %d", len(b), w)
			}
		}
	}
	if err != nil && err != io.EOF { // we expect an EOF
		fmt.Println("\n\nError:", err)
	}

}
