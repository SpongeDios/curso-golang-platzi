package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"Hola", "que", "hace"}

	//El primer valor representa el indice, el segundo el valor
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i := range slice {
		fmt.Println(i)
	}
	//ama = palindrome
	//amor a roma
	palindrome := isPalindrome("AMor a roma")
	fmt.Println(palindrome)
}

func isPalindrome(word string) bool {
	var wordReverse string
	word = strings.ToLower(word)
	for i := len(word) - 1; i >= 0; i-- {
		wordReverse += string(word[i])
	}

	return wordReverse == word
}


