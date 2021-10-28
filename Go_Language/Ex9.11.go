package main

import (
	"fmt"
)

func main(){
var a float64
var b float64
var c float64

fmt.Print("Enter A-: ")
fmt.Scan(&a)

fmt.Print("Enter B- : ")
fmt.Scan(&b)

fmt.Print("Enter C-: ")
fmt.Scan(&c)

data:=CheckSides(a,b,c)

if data==true{
	fmt.Println("These integers could form the sides of a right-angled triangle")
}else {
	fmt.Println("These integers could not form the sides of a right-angled triangle")
}

}

func CheckSides(a float64,b float64,c float64) bool  {
data:=false
if (a==b+c)|| (b==c+a)||(c==a+b) {
	data=true
}
return data
}
