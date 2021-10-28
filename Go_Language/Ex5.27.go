package main

import (
	"fmt"
)

func  main()  {

	var accident int
	var oldpremium float64

	fmt.Print("Enter The Number of accident- ")
	fmt.Scan(&accident)

	fmt.Print("Enter The Old premium- ")
	fmt.Scan(&oldpremium)

	NewPremimum:=PermiumAmount(accident,oldpremium)

	fmt.Println("New Premium is ", NewPremimum)
}


func PermiumAmount(accident int, premiumamount float64) float64{
var newpremium float64

if(accident==0){
newpremium=premiumamount+(premiumamount*2/100)
}else if(accident==1 || accident==2){
	newpremium=premiumamount+(premiumamount*5/100)
}else if(accident==3){
	newpremium=premiumamount+(premiumamount*10/100)
}else if accident>=4 {
	newpremium=premiumamount+(premiumamount*30/100)
}

return newpremium
}