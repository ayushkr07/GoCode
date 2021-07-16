package main

import ("fmt" 
	"math")

func main(){
	var num float64 = 12

	var res = math.Sqrt(num)

	fmt.Printf("res is %.2g",res)
	fmt.Println()

}