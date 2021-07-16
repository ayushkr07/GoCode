package main

import "fmt"

// func add(x int,y int) int {
// 	return x+y
// }

// func main(){
// 	var num1 int=48
// 	var num2 int=7

// 	fmt.Println(add(num1,num2))
// }

// ----------------------------------------------------
// ----------------------------------------------------


// func calc(x int,y int)(int,int){
// 	res1:=x+y
// 	res2:=x-y

// 	return res1,res2
// }

// -----------------------------------------------------

func calc(x int,y int)(res1,res2 int){
	res1 =x+y
	res2 =x-y

	return
}

func main(){
	num1:=23
	num2:=21

	res1,res2:=calc(num1,num2)

	fmt.Println(res1,res2)
}