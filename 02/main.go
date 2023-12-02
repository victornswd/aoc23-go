package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// get if # of color larger than req color
// \w+(?=\s+blue)
var blue = regexp.MustCompile("(\\w+)\\s+blue")
var red = regexp.MustCompile("(\\w+)\\s+red")
var green = regexp.MustCompile("(\\w+)\\s+green")
var number = regexp.MustCompile("[0-9][0-9]?")

var maxBlue = 14
var maxRed = 12
var maxGreen = 13

func getAll(all string, color regexp.Regexp) []int {
	arr := strings.Split(all, ";")
	var nums = []int{}
	for _, i := range arr {
		nums = append(nums, getOne(i, color))
	}
	return nums
}

func getOne(str string, color regexp.Regexp) int {
	a := color.FindString(str)
	a = number.FindString(a)

	x, err := strconv.Atoi(a)
	if err != nil {
		x = 0
	}
	return x
}

func pass(nums []int, color int) bool {
	var bools = []bool{}
	for _, i := range nums {
		bools = append(bools, i <= color)
	}

	bool := slices.Contains(bools, false)

	return !bool
}

func checkAll(all string) int {
	var x = []bool{}
	arr := strings.Split(all, ":")
	newArr := strings.Split(arr[0], "Game ")
	x = append(x, pass(getAll(all, *red), maxRed))
	x = append(x, pass(getAll(all, *blue), maxBlue))
	x = append(x, pass(getAll(all, *green), maxGreen))

	bool := slices.Contains(x, false)

	if !bool {
		x, err := strconv.Atoi(newArr[1])
		if err != nil {
			x = 0
		}
		return x
	}
	return 0
}

func checkMax(all string) int {
	var largestCollection = []int{}
	largestCollection = append(largestCollection, slices.Max(getAll(all, *red)))
	largestCollection = append(largestCollection, slices.Max(getAll(all, *blue)))
	largestCollection = append(largestCollection, slices.Max(getAll(all, *green)))

	var multiply = 1
	for _, i := range largestCollection {
		multiply = multiply * i
	}

	return multiply
}
func main() {
	pwd, _ := os.Getwd()
	d, err := os.ReadFile(pwd + "/02/input.txt")
	if err != nil {
		panic(err)
	}

	str := string(d)
	arr := strings.Split(str, "\n")

	var v = []int{}
	for _, i := range arr {
		v = append(v, checkAll(i))
	}

	var sumP1 = 0
	for _, i := range v {
		sumP1 = sumP1 + i
	}
	fmt.Println("Part 1:", sumP1)

	var vv = []int{}
	for _, i := range arr {
		vv = append(vv, checkMax(i))
	}
	var sumP2 = 0
	for _, i := range vv {
		sumP2 = sumP2 + i
	}
	fmt.Println("Part 2:", sumP2)
}
