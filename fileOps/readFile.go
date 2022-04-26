package main
import ("fmt"
		 "os"
		 "io/ioutil"		
)
func main(){
	if len(os.Args) != 2 {
		fmt.Println("Enter a valid file path")
		return
	}
	fileContent, err := ioutil.ReadFile(os.Args[1])
	if err!= nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(fileContent))
}