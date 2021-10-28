package main

import (
	"fmt"
)

func main(){
	var Number int
	fmt.Print("Enter The Number-: ")
	fmt.Scan(&Number)
	Num:=PlainDrome(Number)

	if(Number==Num){
		fmt.Println("PlainDrome Number")
	}else {
		fmt.Println("Not a PlainDrome Number")
	}
}

func PlainDrome (Number int) int {
	var data int
	var mylist [] int
		i:=5
		for i==5{
			if Number>0{
				Num:=Number%10
				Number=Number/10
				mylist=append(mylist,Num)

			}else {
				i=1
			}
		}

	a:=0
	for i:=0;i< len(mylist);i++{
		if a==0{
			a=1
		}else {
			a=a*10
		}
	}

	NewNum:=0
	for i:=0;i< len(mylist);i++{
	NewNum+=mylist[i]*a
	a=a/10

	}

	data=NewNum

	return data
}