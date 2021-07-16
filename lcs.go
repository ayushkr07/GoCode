package main

import ("fmt"
	"math")

func lcs(x string,y string,n int,m int)float64{
	// fmt.Println(x,y,n,m)
	if n==0 || m==0{
		return 0
	}
	if x[n-1]==y[m-1]{
		return (1+lcs(x,y,n-1,m-1))
	}else{
		res1:=lcs(x,y,n-1,m)
		res2:=lcs(x,y,n,m-1)
		fres:=math.Max(res1,res2)
		return fres
	}

}

func main(){
	x:="samsungelectronic"
	y:="lgelectronic"

	a:=len(x)
	b:=len(y)

	res:=lcs(x,y,a,b)
	res=(res/float64(a+b))*100
	fmt.Println(res)
}