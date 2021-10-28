package main

import (
	"fmt"
)

func main(){
var Year int
fmt.Print("Enter Year- ")
fmt.Scan(&Year)

if ((Year % 4==0) && (Year % 100!=0)) || ((Year %400==0) && (Year % 100!=0)){
	fmt.Print(Year," Is A Leap Year")
}else {
	fmt.Print(Year," Is Not A Leap Year")
}

}