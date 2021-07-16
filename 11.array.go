package main

import "fmt"

func main(){
	// var arr [5]string
	// Output : Five Blank space in array

	// var arr [5]int
	// arr[0]=234
	// arr[4]=987
	// Output : [234 0 0 0 987]

	arr2D:=[2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr2D)

	arr:=[3]int{4,7,8}

	sum:=0

	for i:=0;i<len(arr);i++{
		sum+=arr[i]
	}
	fmt.Println(sum)
}