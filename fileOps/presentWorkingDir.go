package main
import ("fmt"
		"os")
func main(){
	wd, err := os.Getwd()
	if err !=nil{
		fmt.Println(err)
	} 
	fmt.Println(wd)

	if err := os.Chdir("/"); err!=nil{
		fmt.Println(err)
	}
	wd, err = os.Getwd()
	if err !=nil{
		fmt.Println(err)
	} 
	fmt.Println(wd)
}