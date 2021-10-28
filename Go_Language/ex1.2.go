package main

import "fmt"


func main(){
var Hours float64;
var wage float64;

fmt.Print("Enter The Employess Hour- ")
fmt.Scan(&Hours)

fmt.Print("Enter The wage- ")
fmt.Scan(&wage);

FinalRate:=Hours*wage;

fmt.Println("Employee Final wage is ",float64(FinalRate) )
}
