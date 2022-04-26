package main
import ("os"
		"fmt"
		"io"
	)
func main(){
	if len(os.Args)!= 2{
		fmt.Println("Enter a valid file")
		return
	}
	file, err := os.Open(os.Args[1])
	if err!=nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()

	var (
		b = make([]byte, 16)
	)
	for n:=0; err== nil;{
		n, err = file.Read(b);
		if err==nil{
			fmt.Print(string(b[:n]))
		}
	}
	if err != nil && err != io.EOF {
		fmt.Print("\n\nError: ", err)
	}



}