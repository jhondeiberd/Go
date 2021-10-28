package main

import "fmt"

func main(){



	array:=[]int{2,3,0,1,4};

	jump:=0
	steps:=0
	for i:=0;i<=len(array);i=i+jump{
			steps=steps+1
		for j:=i+1;j< len(array);j++{
			 if array[i]<array[j]{
				if jump <(array[i]+array[j]){
					if (jump+array[i]+array[j])<=len(array) {
						jump=jump+array[i]+array[j]
					}
				}

			 }
			 if j==len(array)-1 && jump==0{
				 jump=1
			 }
		}
	}

	fmt.Println("Number of steps to reach end is ",steps)
}

