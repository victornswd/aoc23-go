package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand string
	bid  int
	typ  int
	rank int
}

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

func calcHandType(hand string) int {
	var arr = []int{}
	a := strings.Count(hand, "A")
	k := strings.Count(hand, "K")
	q := strings.Count(hand, "Q")
	j := strings.Count(hand, "J")
	t := strings.Count(hand, "T")
	nine := strings.Count(hand, "9")
	eig := strings.Count(hand, "8")
	sev := strings.Count(hand, "7")
	six := strings.Count(hand, "6")
	fiv := strings.Count(hand, "5")
	fur := strings.Count(hand, "4")
	thr := strings.Count(hand, "3")
	two := strings.Count(hand, "2")
	arr = append(arr, a, k, q, j, t, nine, eig, sev, six, fiv, fur, thr, two)

	if slices.Contains(arr, 5) {
		return 7
	}
	if slices.Contains(arr, 4) {
		return 6
	}
	if slices.Contains(arr, 3) {
		if slices.Contains(arr, 2) {
			return 5
		}
		return 4
	}
	if slices.Contains(arr, 2) {
		x := 0
		for _, v := range arr {
			if v == 2 {
				x++
			}
		}
		if x == 2 {
			return 3
		}
		return 2
	}
	if slices.Contains(arr, 1) {
		return 1
	}

	return 0
}

func calcHandTypeP2(hand Hand) int {
	if !strings.Contains(hand.hand, "J") {
		return calcHandType(hand.hand)
	}

	pos := strings.Index(hand.hand, "J")
	prev := 0
	next := 0
	if pos-1 >= 0 {
		prev = calcHandType(strings.ReplaceAll(hand.hand, "J", string(hand.hand[pos-1])))
	}

	if pos+1 < 5 {
		lastJ := strings.LastIndex(hand.hand, "J")
		if lastJ > pos+1 && lastJ < 4 {
			next = calcHandType(strings.ReplaceAll(hand.hand, "J", string(hand.hand[lastJ+1])))
		} else {
			next = calcHandType(strings.ReplaceAll(hand.hand, "J", string(hand.hand[pos+1])))
		}
	}

	if prev >= next {
		return prev
	} else if next > prev {
		return next
	}

	return 0
}

func main() {
	pwd, _ := os.Getwd()
	var arr = []string{}

	d, err := os.ReadFile(pwd + "/07/input.txt")
	if err != nil {
		panic(err)
	}
	str := string(d)
	arr = strings.Split(str, "\n")

	cards := make(map[string]int)
	cards["A"] = 14
	cards["K"] = 13
	cards["Q"] = 12
	cards["J"] = 11
	cards["T"] = 10
	cards["9"] = 9
	cards["8"] = 8
	cards["7"] = 7
	cards["6"] = 6
	cards["5"] = 5
	cards["4"] = 4
	cards["3"] = 3
	cards["2"] = 2

	var hands = []Hand{}
	for _, v := range arr {
		hand := strings.Split(v, " ")

		if len(hand) > 1 {
			hnd := new(Hand)
			hnd.hand = hand[0]
			hnd.bid = strToIntSimp(hand[1])
			hands = append(hands, *hnd)
		}
	}

	for i, v := range hands {
		v.typ = calcHandType(v.hand)
		hands[i] = v
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].typ == hands[j].typ {
			x := strings.SplitAfterN(hands[i].hand, "", -1)
			y := strings.SplitAfterN(hands[j].hand, "", -1)

			for i := 0; i < len(x); i++ {
				if x[i] != y[i] {
					if cards[x[i]] > cards[y[i]] {
						return false
					} else {
						return true
					}
				}
			}
		}

		return hands[i].typ < hands[j].typ
	})

	rankP1 := 0
	for i := 0; i < len(hands); i++ {
		rankP1++
		hands[i].rank = rankP1
	}

	var winningsP1 int
	for _, v := range hands {
		winningsP1 += (v.bid * v.rank)
	}

	fmt.Println("Part 1:", winningsP1)

	jokers := cards
	jokers["J"] = 1

	var handsP2 = []Hand{}
	for _, v := range arr {
		hand := strings.Split(v, " ")

		if len(hand) > 1 {
			hnd := new(Hand)
			hnd.hand = hand[0]
			hnd.bid = strToIntSimp(hand[1])
			handsP2 = append(handsP2, *hnd)
		}
	}
	for i, v := range handsP2 {
		v.typ = calcHandTypeP2(v)
		handsP2[i] = v
	}

	sort.Slice(handsP2, func(i, j int) bool {
		if handsP2[i].typ == handsP2[j].typ {
			x := strings.SplitAfterN(handsP2[i].hand, "", -1)
			y := strings.SplitAfterN(handsP2[j].hand, "", -1)

			for i := 0; i < len(x); i++ {
				if x[i] != y[i] {
					if jokers[x[i]] > jokers[y[i]] {
						return false
					} else {
						return true
					}
				}
			}
		}

		return handsP2[i].typ < handsP2[j].typ
	})

	rankP2 := 0
	for i := 0; i < len(handsP2); i++ {
		rankP2++
		handsP2[i].rank = rankP2
	}
	// fmt.Println(handsP2)

	var winningsP2 int
	for _, v := range handsP2 {
		winningsP2 += (v.bid * v.rank)
	}

	fmt.Println(winningsP2)
}
