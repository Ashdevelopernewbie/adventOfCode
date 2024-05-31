package main

func checkFirstAndLastIndex(word string) (bool, string) {
	var firstChar rune
	var lastChar rune
	finalNum := ""

	for _, char := range word[:3] {
		if '0' <= char && char <= '9' {
			firstChar = char
			break
		}
	}

	for _, char := range word[len(word)-3:] {
		if '0' <= char && char <= '9' {
			lastChar = char
			break
		}
	}

	if firstChar == 0 || lastChar == 0 {
		return false, ""
	}
	finalNum = string(firstChar) + string(lastChar)
	return true, finalNum
}
