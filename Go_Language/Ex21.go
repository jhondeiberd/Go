package main

import (
	"fmt"
)

func main(){

	nums :=[]int {1,5,9,1,5,9}
	k:= 2
	t:= 3
Content:=CheckCondition(nums,k,t)
fmt.Println("Output is",Content)


}


func CheckCondition(body [] int, jump int, target int) bool{

for  i:=0;i< len(body);i++{
	for j:=i+jump;j<len(body);j++{
		data:=(body[i]-body[j])
		if data <0{

				data=data*-1
			}
		if data <=target{
			return true
		}else {
			j=len(body)
		}
		}

		}



	return false
}