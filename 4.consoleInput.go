package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	// scanner:=bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type Something : ")
	// scanner.Scan()
	// input:=scanner.Text()
	// fmt.Printf("You typed %q",input)
	// fmt.Println()

	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Printf("Type Something : ")
	scanner.Scan()
	input,_:=strconv.ParseInt(scanner.Text(),10,64)
	fmt.Printf("%d %T",input,input)
}