package main
/*
Given two strings word1 and word2 randomly, return the maximum number of operations required to convert word1 to word2.
You have the following three operations permitted on a word:
Insert a character
Delete a character
Replace a character

Restriction: the maximum number of times using one of the permitted operations can NOT be more than the number of characters of word1 OR word2.
Case: word1 has 4 characters. Word2 has 6 characters. You can’t use any operation more than 6 times.
 */
import (
	"fmt"
	"strings"
)

func main() {
	//var word1 string = "horse"
	var word1 string = "abcd"
	var word2 string = "rosddd"
	var word3 string = "intention"
	var word4 string = "execution"
	ChangeCharactWord(word1, word2)
	ChangeCharactWord(word3, word4)

}

func ConvertStringtoArray (word string) []string {
	var arr[]string
	for i:=0; i< len(word); i++ {
		arr = append(arr, string(word[i]))
	}
	return arr
}

func ChangeCharactWord (word1 string, word2 string){
	var arrWord1 []string = []string(ConvertStringtoArray(word1))
	var arrWord2 []string = []string(ConvertStringtoArray(word2))
	var lenWord1 int = len(arrWord1)
	var lenWord2 int = len(arrWord2)
	var first, second, newWord1, newWord2 []string

	if lenWord1 >= lenWord2{
		first = arrWord1
		second = arrWord2
	} else{
		first = arrWord2
		second = arrWord1
	}

	for i:= 0; i < len(second); i++{
		newWord1 = append(newWord1, second[i])
		newWord2 = append(newWord2, first[i])
	}

	for i:= len(second); i < len(first); i++{
		newWord2 = append(newWord2, first[i])
	}
	fmt.Println("\nThe words to change is:\t\tWord1 = ", word1, "\t\tand\t\tWord2 = ", word2)
	if lenWord1 >= lenWord2{
		fmt.Println("The words changed is:\t\tWord1 = ", strings.Join(newWord1,""), "\t\tand\t\tWord2 = ", strings.Join(newWord2,""))
	} else {
		fmt.Println("The words changed is:\t\tWord1 = ", strings.Join(newWord2,""), "\t\tand\t\tWord2 = ", strings.Join(newWord1,""))
	}

}