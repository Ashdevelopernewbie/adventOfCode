package main

import (
	"fmt"
)

func checkFirstAndLastIndex(word string) (bool, string) {
	var firstNum, lastNum string
	firstChar := string(word[0])
	lastChar := string(word[len(word)-1])

	switch {
	case firstChar == "1":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	case firstChar == "2":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	case firstChar == "3":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	case firstChar == "4":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	case firstChar == "5":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	case firstChar == "6":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	case firstChar == "7":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	case firstChar == "8":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	case firstChar == "9":
		fmt.Printf("%s, found\n", firstChar)
		firstNum = firstChar
	default:
		fmt.Println("No numbers found in first index")
		return false, ""
	}

	switch {
	case lastChar == "1":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	case lastChar == "2":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	case lastChar == "3":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	case lastChar == "4":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	case lastChar == "5":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	case lastChar == "6":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	case lastChar == "7":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	case lastChar == "8":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	case lastChar == "9":
		fmt.Printf("%s, found\n", lastChar)
		lastNum = lastChar
	default:
		fmt.Println("No numbers found in last index")
		return false, ""
	}

	finalNum := firstNum + lastNum
	return true, finalNum
}
