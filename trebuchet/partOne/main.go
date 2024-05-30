package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var validNumbers = map[string]string{
	"1":  "one",
	"2":  "two",
	"3":  "three",
	"4":  "four",
	"5":  "five",
	"6":  "six",
	"7":  "seven",
	"8":  "eight",
	"9":  "nine",
	"10": "ten",
}

// KEEP IT STUPID SIMPLE
var sum int

func sumCalibrationValue(value int) int {
	sum += value
	return sum
}

func calibrationValueFromWords(words string) int {
	// Find the numbers in the sample data
	numbers := []rune{}
	for _, r := range words {
		if _, ok := validNumbers[string(r)]; ok {
			numbers = append(numbers, r)
		}
	}

	// Find first and last digit from calibration values
	firstDigit := numbers[0]
	lastDigit := numbers[len(numbers)-1]

	// Convert to integer
	valueStr := fmt.Sprintf("%c%c", firstDigit, lastDigit)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(err)
	}

	return sumCalibrationValue(value)
}

func main() {
	sampleFile, err := os.Open("final.base")
	check(err)
	defer sampleFile.Close()

	scanner := bufio.NewScanner(sampleFile)
	for scanner.Scan() {
		calibrationValue := calibrationValueFromWords(scanner.Text())
		fmt.Println(calibrationValue)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// Cleaned up code by standardizing variable names, removing debugging statements,
// improving readability, and more. Changed filename and scanner variables to better
// reflect their purpose. Removed unnecessary comments. Used check function to
// handle errors in a consistent way.
