package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var number = regexp.MustCompile("[0-9][0-9]?[0-9]?[0-9]?")

func strToIntSimp(str string) int {
	x, err := strconv.Atoi(str)
	if err != nil {
		x = 0
	}
	return x
}
func strToInt(str []string) []int {
	var arr = []int{}
	for _, i := range str {
		x, err := strconv.Atoi(i)
		if err != nil {
			arr = append(arr, 0)
		}
		arr = append(arr, x)
	}
	return arr
}

func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}

func calcPoints(nums []int) int {
	a := len(nums)

	point := int(math.Pow(2, float64(a)-1))
	return point
}

func getTime(str string) []int {
	a := strings.Split(str, "Time: ")
	b := number.FindAllString(a[1], -1)
	return strToInt(b)
}

func getDist(str string) []int {
	a := strings.Split(str, "Distance: ")
	b := number.FindAllString(a[1], -1)
	return strToInt(b)
}
func getTimeNoSpace(str string) []int {
	a := strings.Split(str, "Time: ")
	b := number.FindAllString(a[1], -1)
	var c string
	for _, v := range b {
		c = c + v
	}
	var d []int
	d = append(d, strToIntSimp(c))

	return d
}

func getDistNoSpace(str string) []int {
	a := strings.Split(str, "Distance: ")
	b := number.FindAllString(a[1], -1)
	var c string
	for _, v := range b {
		c = c + v
	}
	var d []int
	d = append(d, strToIntSimp(c))

	return d
}
func maths(time int, distance int) int {

	var total []int
	var records []int
	for i := 0; i < time; i++ {
		total = append(total, (i * (time - i)))
	}

	for _, v := range total {
		if v > distance {
			records = append(records, v)
		}
	}

	return len(records)
}

func main() {
	pwd, _ := os.Getwd()
	var arr = []string{}

	d, err := os.ReadFile(pwd + "/06/input.txt")
	if err != nil {
		panic(err)
	}
	str := string(d)
	arr = strings.Split(str, "\n")

	time := getTime(arr[0])
	distance := getDist(arr[1])

	var maps = make(map[int]int)
	for i := 0; i < len(time); i++ {
		maps[time[i]] = distance[i]
	}

	var winnings []int
	for key, value := range maps {
		winnings = append(winnings, maths(key, value))
	}

	var sumP1 = 1
	for _, i := range winnings {
		sumP1 *= i
	}

	fmt.Println("Part 1:", sumP1)

	time2 := getTimeNoSpace(arr[0])
	distance2 := getDistNoSpace(arr[1])

	var maps2 = make(map[int]int)
	for i := 0; i < len(time2); i++ {
		maps2[time2[i]] = distance2[i]
	}

	var winnings2 []int
	for key, value := range maps2 {
		winnings2 = append(winnings2, maths(key, value))
	}

	var sumP2 = 1
	for _, i := range winnings2 {
		sumP2 *= i
	}
	fmt.Println("Part 2:", sumP2)
}
