package main

import "fmt"

func main(){
	var arra[3] int
	arra[0]=1
	arra[1]=2
	arra[2]=3

	arra[0],arra[1]=arra[1],arra[0]

	fmt.Println(arra[0])
	fmt.Println(arra[1])

}
