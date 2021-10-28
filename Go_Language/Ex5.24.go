package main

import "fmt"

func main(){
var percentage float64
data:=grade(percentage)
fmt.Print(data)
}


func grade(percentage float64) string {
	var letter string



	i:=5
	for i==5{
		fmt.Print("Enter the percentage- ");
		fmt.Scan(&percentage)
		if percentage>100 || percentage <=0{
			fmt.Println("Invalid Input")
			fmt.Println("Try Again")
			i=5
		} else {
			if percentage>=90 && percentage<=100{
				letter="A";
				i=1
			}
			if percentage>=80 && percentage<=89{
				letter="B";
				i=1
			}

			if percentage>=70 && percentage<=79{
				letter="C";
				i=1
			}

			if percentage>=60 && percentage<=69{
				letter="D";
				i=1
			}
			if percentage<60 {
				letter="F";
				i=1
			}

		}
	}

	return letter
}