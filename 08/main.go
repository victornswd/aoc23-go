package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var letters = regexp.MustCompile("[A-Z][A-Z][A-Z]")

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

type Node struct {
	Name       string
	Directions []string
}

func theLoop(dirArr []int, nodes []Node) {
	for i := 0; i < len(dirArr); i++ {
		steps++

		if nextNode.Name == "" {
			for _, v := range nodes {
				if v.Name == "AAA" {
					nextNode = v
				}
			}
		}
		next = nextNode.Directions[dirArr[i]]

		for _, v := range nodes {
			if v.Name == next {
				nextNode = v
			}
		}

		if nextNode.Name == "ZZZ" {
			break
		}

		if nextNode.Name != "ZZZ" && i == len(dirArr)-1 {
			theLoop(dirArr, nodes)
		}
	}
}

var steps = 0
var next = ""
var nextNode = Node{}

func main() {
	pwd, _ := os.Getwd()
	var arr = []string{}

	d, err := os.ReadFile(pwd + "/08/input.txt")
	if err != nil {
		panic(err)
	}
	str := string(d)
	arr = strings.Split(str, "\n")

	directions := arr[0]
	directions = strings.ReplaceAll(directions, "L", "0")
	directions = strings.ReplaceAll(directions, "R", "1")
	dirArr := strToInt(strings.SplitAfterN(directions, "", -1))

	var nodes = []Node{}
	for _, i := range arr[2:] {
		if len(i) == 0 {
			continue
		}

		name := strings.Split(i, " = ")[0]
		dir := letters.FindAllString(i, -1)
		var arr = []string{}
		arr = append(arr, dir[1])
		arr = append(arr, dir[2])
		node := new(Node)
		node.Name = name
		node.Directions = arr

		nodes = append(nodes, *node)
	}

	theLoop(dirArr, nodes)
	fmt.Println("Part 1:", steps)
}
