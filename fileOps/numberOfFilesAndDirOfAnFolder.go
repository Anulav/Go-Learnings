package main
import ("fmt"
		"os"
		"path/filepath"
	)
func main(){
	if len(os.Args)!=2 {
		fmt.Println("Enter a path to get Absoulte Path")
		return
	}
	root, err := filepath.Abs(os.Args[1])
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(root)
	 var c struct{
		 files int
		 directories int
	 }
	 filepath.Walk(root, func(path string,info os.FileInfo,err error) error{
		if info.IsDir() {
			c.directories++
		}else{
			c.files++
		}
		fmt.Println("-",path)
		return nil
	 })
	 fmt.Printf("The folder contains %d files and %d directories ",c.files,c.directories)
}