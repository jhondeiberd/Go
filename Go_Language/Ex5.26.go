package main

import (
	"fmt"
)

func main(){
var size1 float64
var size2 float64
var size3 float64

data:=check2(size1,size2,size3)

fmt.Println(data)
}

func check2(size1 float64, size2 float64, size3 float64) string {
var datas string

fmt.Print("Enter size1- ")
fmt.Scan(&size1)

fmt.Print("Enter Size2 -")
fmt.Scan(&size2)

fmt.Print("Enter Size3- ")
fmt.Scan(&size3)

if(size1==size2 && size1==size3){
	datas="Equilateral triangle"
} else if(size1==size2 || size2==size3|| size1==size3){
	datas="Isosceles triangle "
}else {
	datas="Scalene triangle"
}

return datas
}
