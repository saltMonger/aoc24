package puzzle1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"tfx-aoc24/utils"
)

var list1 = []int{}
var list2 = []int{}

func readInput(path string) {
	f, err := os.Open(path)
	utils.Check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := strings.Split(strings.ReplaceAll(scanner.Text(), "   ", " "), " ")
		if len(value) != 2 {
			panic(fmt.Errorf("invalid list splitting"))
		}

		l1, err := strconv.Atoi(value[0])
		utils.Check(err)
		l2, err := strconv.Atoi(value[1])
		utils.Check(err)
		list1 = append(list1, l1)
		list2 = append(list2, l2)
	}
}

func Puzzle(path string) {
	readInput(path)
	sum := 0
	// sort
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	// diff
	for i := range list1 {
		value := list1[i] - list2[i]
		if value < 0 {
			value = -value
		}
		fmt.Println("dist", i, ":", value)
		sum += value
	}

	fmt.Println("Puzzle 1 distance:", sum)
}

// mmm brute
func enumerateList2Hits(search int) int {
	hits := 0
	for _, x := range list2 {
		if x == search {
			hits++
		}
	}
	return hits
}

func PuzzleHard(path string) {
	readInput(path)
	simScore := 0

	simMap := map[int]int{}

	for _, x := range list1 {
		if _, ok := simMap[x]; !ok {
			simMap[x] = enumerateList2Hits(x) * x
			simScore += simMap[x]
		} else {
			simScore += simMap[x]
		}
	}

	fmt.Println("Puzzle 1 sim score: ", simScore)
}
