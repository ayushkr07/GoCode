package main

import "fmt"

func main(){
	// i:=1
	// for i<=5{
	// 	fmt.Println("Hello",i)
	// 	i+=1
	// }

	for i:=1;i<=100;i+=1{
		if(i == 5){
			continue
		}
		fmt.Println(i)
	}
}