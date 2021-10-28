package main

import (
	"fmt"
)

func main(){
var data int
	fmt.Print("Enter 4 digits Number-:  ")
	datas:=Validation(data)
	Data:=Encryption(datas)
	fmt.Println("Encrypted data is- ",Data)
	Dec:=Decryption(Data)
	fmt.Println("Decrupted data is- ",Dec)
}
func Decryption(data [] int)[] int{
var Decrupt [] int

for i:=0;i< len(data);i++{
	Decrupt=append(Decrupt,((data[i]+3)%10))
}

	return Decrupt
}


func Encryption(data int) [] int {
	var Arrays [] int


	for data>0{
		NewNum:=data%10
		data=data/10
		NewNum=(NewNum+7)%10
		Arrays=append(Arrays,NewNum)

	}


	return Arrays
}




func Validation(data int) int {
	i:=5
	for i==5 {

		_, err := fmt.Scanln(&data)

		if err != nil {
			fmt.Println("No String")
			fmt.Print("Try Again -: ")
			i = 5
		}else if data>9999 || data<=999{
			fmt.Println("Number Should be between 1000 and 9999")
			fmt.Print("Try Again -: ")
			i=5
		}else {
			i=1
		}


	}
	return data
}



