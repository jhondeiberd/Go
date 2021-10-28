package main

import "fmt"

func main(){
	var day int
	var month int
	var year int

	fmt.Print("Enter day-: ")
	fmt.Scan(&day)

	fmt.Print("Enter Month-: ")
	fmt.Scan(&month)

	fmt.Print("Enter Year -: ")
	fmt.Scan(&year)

	if month==1 || month==3 || month==5|| month==7 || month==8|| month==10|| month==12{
		if day==31 && month==12{
			day=1
			month=1
			year=year+1
		}else if day==31 && month!=12 {
			day=1
			month=month+1
		}else {
			day=day+1
		}


	}else if month==4 || month==6|| month==9|| month==11{
		if day==30{
			day=1
			month=month+1
		}else {
			day=day+1
		}
	}else if month==2 {
		if day==29{
			day=1
			month=month+1
		}else {
			day=day+1

		}
	}

fmt.Println("Date is",day,"-",month,"-",year," ")

}
