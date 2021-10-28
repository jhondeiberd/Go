package main

import (
	"fmt"
	"strings"
)

func main(){


	board:= [] [] string {{"A","B","C","E"},{"S","F","C","S"},{"A","D","E","E"} };
	word:="ABCCED"
	Data:=TableChecker(board,word)
	fmt.Println(Data)

}


func TableChecker(word [][] string, data string) [][] int{
var NewArray [] [] int




	var Status [5] [5] bool


	for i:=0;i<len(word);i++{

		for j:=0;j<len(word[i]);	j++{

			Status[i][j]=false
		}

	}


	NewData:=wordtoarray(data)
	count:=0

		for i:=0;i<len(word);i++{
			for j:=0;j<len(word[i]);j++{
				if count==len(NewData){
					break
				}
				if word[i][j]==NewData[count] && Status[i][j]==false{
					row1:=[] int{i,j}

						NewArray=append(NewArray,row1)
						count=count+1
						Status[i][j]=true


				}
				if i==len(word)-1 && j==len(word[0])-1 {
					i=0
					j=0
				}


			}
		}

	return NewArray
}


func wordtoarray(word string) [] string {
	NewArray:=[] string {}

	for i:=0;i<len(word);i++{
		split:= strings.Split(word,"")
		NewArray=append(NewArray,split[i])
	}

	return NewArray
}