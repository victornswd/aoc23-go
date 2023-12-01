package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func combineNumbers(nums []string) int {
	first := nums[0]
	last := nums[len(nums)-1]

	num, err := strconv.Atoi(first + last)
	if err != nil {
		panic(err)
	}
	return num
}

func sumOfCalibration(nums []int) int {
	var sum = 0
	for _, i := range nums {
		sum += i
	}

	return sum
}

func allNumbersArray(z string) []string {
	nums := regexp.MustCompile("[0-9]")
	numArr := nums.FindAllString(z, -1)

	return numArr
}

func convertStrToNum(nums string) string {
	var converted = ""
	converted = strings.ReplaceAll(nums, "one", "o1e")
	converted = strings.ReplaceAll(converted, "two", "t2o")
	converted = strings.ReplaceAll(converted, "three", "t3e")
	converted = strings.ReplaceAll(converted, "four", "f4r")
	converted = strings.ReplaceAll(converted, "five", "f5e")
	converted = strings.ReplaceAll(converted, "six", "s6x")
	converted = strings.ReplaceAll(converted, "seven", "s7n")
	converted = strings.ReplaceAll(converted, "eight", "e8t")
	converted = strings.ReplaceAll(converted, "nine", "n9e")

	return converted
}

func main() {
	d, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	str := string(d)
	arr := strings.Fields(str)

	var calibrationVal = []int{}
	for _, i := range arr {
		calibrationVal = append(calibrationVal, combineNumbers(allNumbersArray(i)))
	}

	fmt.Println("Part 1:", sumOfCalibration(calibrationVal))

	var newVal = []int{}
	for _, i := range arr {
		newVal = append(newVal, combineNumbers(allNumbersArray(convertStrToNum(i))))
	}

	fmt.Println("Part 2:", sumOfCalibration(newVal))
}
