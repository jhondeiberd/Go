package main

import "fmt"

func main ()  {

	var year1 int
	var year2 int
	var mon1 int
	var mon2 int
	var day1 int
	var day2 int

	fmt.Print("Enter Year 1- ")
	fmt.Scan(&year1)

	fmt.Print("Enter Year 2- ")
	fmt.Scan(&year2)

	fmt.Print("Enter Month 1- ")
	fmt.Scan(&mon1)

	fmt.Print("Enter Month 2- ")
	fmt.Scan(&mon2)

	fmt.Print("Enter Day1- ")
	fmt.Scan(&day1)

	fmt.Print("Enter Day2- ")
	fmt.Scan(&day2)

	status:=check(year1,year2,mon1,mon2,day1,day2);

	if status == true{
		fmt.Println("year1 ",year1,"-",mon1,"-",day1 ," Comes first")
	}else {
		fmt.Println("year2 ",year2,"-",mon2,"-",day2, " Comes first")
	}

}
func check(year1 int ,year2 int ,mon1 int,mon2 int,day1 int,day2 int) bool {


	if year1 < year2{
		return true
	}else {
		return false
	}

	if year1==year2 {
		if mon1<mon2{
			return true
		}
	}else {
		return false
	}

	if year1==year2{
		if mon1==mon2{
			if day2<day2{
				return true
			}
		}
	}else {
		return false
	}



	return false
}