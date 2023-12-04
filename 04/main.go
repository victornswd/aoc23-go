package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var number = regexp.MustCompile("[0-9][0-9]?")
var delimiter = regexp.MustCompile("[:|]")

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

func splitLine(str string) []int {
	twoSets := delimiter.Split(str, -1)

	if len(twoSets) > 0 {
		twoSets = twoSets[1:]

	}
	if len(twoSets) > 0 {
		win := number.FindAllString(twoSets[0], -1)
		all := number.FindAllString(twoSets[1], -1)
		hash := HashGeneric(strToInt(win), strToInt(all))
		return hash
	}

	return []int{}
}

func calcPoints(nums []int) int {
	a := len(nums)

	point := int(math.Pow(2, float64(a)-1))
	return point
}

func main() {
	pwd, _ := os.Getwd()
	var arr = []string{}

	d, err := os.ReadFile(pwd + "/04/input.txt")
	if err != nil {
		panic(err)
	}
	str := string(d)
	arr = strings.Split(str, "\n")

	var vv = []int{}
	for _, i := range arr {
		if len(i) == 0 {
			continue
		}
		vv = append(vv, calcPoints(splitLine(i)))
	}

	var sumP1 = 0
	for _, i := range vv {
		sumP1 += i
	}

	fmt.Println("Part 1:", sumP1)

	cards := make(map[int]int)
	for id, i := range arr {
		if len(i) == 0 {
			continue
		}
		cardNum := id + 1
		cards[cardNum]++
		matches := len(splitLine(i))

		for i := 1; i <= matches; i++ {
			cards[cardNum+i] += cards[cardNum]
		}
	}

	var sumP2 = 0
	for _, v := range cards {
		sumP2 += v
	}

	fmt.Println("Part 2:", sumP2)
}
