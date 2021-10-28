package main

import (
	"fmt"
)

func main()  {
var Scores[6] int

Avergage:=Scoresss(Scores)

fmt.Println("Average Score is ",Avergage)
}

func Scoresss(Score [6]int ) float64{
var Average float64
Average=0

for i:=0;i< len(Score);i++{
	fmt.Print("Enter Score ",i," -: ")
	fmt.Scan(&Score[i])

}
 min:=Score[0]
 max:=0

 for i:=0;i< len(Score);i++{
	 if Score[i] >max{
		 max=Score[i]
	 }

	 if(min>Score[i]){
		 min=Score[i]
	 }
 }

	data:=findAndDelete(Score,max,min)

	for i:=0;i< len(data);i++{
		Average+=float64(data[i])
	}
	Average=Average/ float64(len(data))

	return Average
}

func findAndDelete(s [6]int, max int,min int) []int {
	index := 0
	for i:=0;i< len(s);i++ {
		if s[i] != max && s[i] !=min {
			s[index] = s[i]
			index++
		}



	}
	return s[:index]
}