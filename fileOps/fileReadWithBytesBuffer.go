package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b = bytes.NewBuffer(make([]byte, 26))
	//
	var texts = []string{
		`Johnny Depp is innocent`,
		`Some women are psychopaths`,
	}
	for i := range texts {
		b.Reset()
		b.WriteString(texts[i])
		fmt.Println("Length ", b.Len(), "\nCapacity ", b.Cap())
	}
}
