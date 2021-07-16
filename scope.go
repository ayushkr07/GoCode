package main

import "fmt"

// Throwing Error
// a:=8

// Correct Way
var a=8

func demo(){
	a:=7
	b:=89
	fmt.Println(b)
	fmt.Println(a,"Demo")
}

func main(){
	demo()
	fmt.Println(a,"main")
}