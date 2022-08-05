package main

import (
	"fmt"
	"math"
	"strconv"
)

func stringToInt(str string) int {
	// convert to int
	convertedStr, err := strconv.Atoi(str)

	if err != nil {
		panic(fmt.Sprintf("Failed to convert string to int: %v", err))
	}

	return convertedStr
}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision
	convertedStr, err := strconv.ParseFloat(str, 64)

	if err != nil {
		panic(fmt.Sprintf("Failed to convert string to float: %v", err))
	}

	return math.Floor(convertedStr*10000) / 10000
}

func FloatToString(value float64) string {
	// convert to float with 2 digits of precision

	str := strconv.FormatFloat(value, 'f', 2, 64)

	return str
}

func TestStringConverstions() {
	isFailed := false
	if stringToInt("10") != 10 {
		fmt.Println("Failed: stringToInt")
		isFailed = true
	}

	if stringToFloat("123.33333333333") != 123.3333 {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if FloatToString(1.0/3) != "0.33" {
		fmt.Println("Failed: floatToString")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}
