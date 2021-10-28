package main

import (
	"fmt"
	"strings"
)

func main(){

	 s := "barfoofoobarthefoobarman"
	words:= []string{"bar","foo","the"}

	Data:=FinalAlgo(s,words)
	fmt.Println(Data)
}

func FinalAlgo(Word string, WordArray [] string)[] int{
var data [] int
StrToArr:=StringToArray(Word)


AllPosss:=AllPossiblities(WordArray)
var Conce [] string
for i:=0;i< len(AllPosss);i++{
	Value:=MultiToSingle(AllPosss[i])

	Conce=append(Conce,Value)

}

	for i:=0;i< len(Conce);i++{
	Words:=StringToArray(Conce[i])
	count:=0
	for j:=0;j< len(StrToArr);j++{
		if Words[count]==StrToArr[j]{


			count=count+1

			if count==len(Words)-1{
				val:=(len(Words)-2)-j
				if val<0{
					val=val *-1
				}
				data=append(data,val)
				break
			}
		}else {

			if count>0{
				count=0
				j=j-1
			}
			count=0

		}
	}
}

	return data
}

func CheckValues(){

}

func AllPossiblities(WordArray [] string)[][] string{
	var Value [] []string
	n:=len(WordArray)
	var backTrack func(int)
	backTrack=func(first int) {
		if first == n {
			temp:=make([]string, n)
			copy(temp,WordArray)
			Value = append(Value, temp)
		}
		for i:=first;i<n;i++ {
			WordArray[first],WordArray[i] = WordArray[i],WordArray[first]
			backTrack(first+1)
			WordArray[first],WordArray[i] = WordArray[i],WordArray[first]
		}

	}



	backTrack(0)
return Value
}



func StringToArray(Word string)[] string{
	var data [] string
	s:=strings.Split(Word,"")
	for i:=0;i<len(Word);i++{

		data=append(data,s[i])
	}

	return data
}

func MultiToSingle(value [] string) string{
var data  string
d:=""
for i:=0;i< len(value);i++{
d=d+value[i]
}
data=d
	return data
}

