package main

import "fmt"

func main(){
	var num int
	fmt.Print("Enter The Number-: ")
	fmt.Scan(&num)
	Data:=fctorial(num)
	fmt.Println(Data)
}

func fctorial(num int) int{

	data:=1
	for i:=1;i<=num;i++{
		data=data*i
	}
	return data
}