package main

import (
	"fmt"
	"strings"
)

/* Exercise1
Given an input string s, reverse the order of the words. A word is defined as a sequence of non-space characters.
The words in s will be separated by at least one space. Return a string of the words in reverse order concatenated by a single space.
Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only
have a single space separating the words. Do not include any extra spaces.
 */

/* Exercise3

 */

func main() {
	fmt.Println("\n**** Exercise 1  *****")
	var word1, word2, word3, word4, word5 string
	word1 = "the sky is blue"
	word2 = "  hello world  "
	word3 = "a good   example"
	word4 = "      Bob               Loves                       Alice   "
	word5 = "Alice does not even like bob"
	//fmt.Println(SliceWord(word1))
	SliceWord(word1)
	fmt.Println(SliceWord(word2))
	fmt.Println(SliceWord(word3))
	fmt.Println(SliceWord(word4))
	fmt.Println(SliceWord(word5))
}

func SliceWord (word string) string {
	cleanWord := strings.Fields(word)
	fmt.Println(cleanWord)
	reversePosition(&cleanWord, 0, len(cleanWord)-1)
	fmt.Println(cleanWord)
	return strings.Join(cleanWord, " ")
}

func reversePosition (m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		fmt.Println(m,"/...", (*m)[i],"/", (*m)[j],"/==", (*m)[j],"/",(*m)[i])
		i++
		j--
	}
}

