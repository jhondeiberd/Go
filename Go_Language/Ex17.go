package main

import (
	"fmt"
)

func main (){

data:=[]int{1,0}
Final:=Arraytoint(data)

fmt.Println("Number To Array is ",Final);
}


func Arraytoint(array [] int) int{
data :=1
increase:=1

for i:=0;i< len(array);i++{
if increase==1{
	data=array[i]*increase
	increase=10
}else {
		if len(array)==2{
			if array[i]==0{
				data=data*increase
				increase=increase*10

			}
		}else{
			data=data+(array[i]*increase)
			increase=increase*10
		}
		}

}

	return data
}