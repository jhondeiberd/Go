package main

import "fmt"

func main(){
	num:=1

	for i:=1;i<=15;i++{
		if(i%2!=0){
			num=num*i
		}
	}

	fmt.Println(num);
}