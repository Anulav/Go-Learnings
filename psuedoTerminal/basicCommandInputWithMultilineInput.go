package main
func main(){

}

func ScanArgs(data []byte, atEOF bool) (advance int, token []byte, err error){
	start, first := 0, rune(0)
	for width :=0; start <len(data); start+=width {
		first, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(first){
			break
		}
		if isQuote(first){
			start++
		}
	}
}