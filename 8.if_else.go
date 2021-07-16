package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your age : ")
	scanner.Scan()
	age,_ := strconv.ParseInt(scanner.Text(),10,64)

	if(age <= 11){
		fmt.Println("You are child")
	}else if(age > 11 && age <=19){
		fmt.Println("You are teenager")
	}else{
		fmt.Println("You are adult")
	}
	

}