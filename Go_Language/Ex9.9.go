package main

import (
	"fmt"
)

func main()  {
	var Decimal int
	fmt.Print("Enter The Decimal ")
	fmt.Scan(&Decimal)
	Data:=DecimalToBinary(Decimal)

	fmt.Println("Binary to Decimal for ",Decimal," is",Data)
}

func DecimalToBinary(Number int) int {
	var Binary int
	Binary=0
	var Arrays [] int
		for Number>0{
			NewNum:=Number%2
			Arrays=append(Arrays,NewNum)
			Number=Number/2
		}

		a:=0
	for i:=0;i< len(Arrays);i++{
	if a==0{
		Binary=Binary+Arrays[i]
		a=10
	}else {
		Binary=Binary+(Arrays[i]*a)
		a=a*10
	}

	}

	return Binary
}