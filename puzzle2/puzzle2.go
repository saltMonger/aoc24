package puzzle2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tfx-aoc24/utils"
)

type reactorLevel []int

func (r reactorLevel) IsReactorLevelDecreasing() bool {
	var last int
	for i, x := range r {
		if i == 0 {
			last = x
			continue
		}
		if x > last {
			return false
		}
		last = x
	}
	return true
}

func (r reactorLevel) IsReactorLevelIncreasing() bool {
	var last int
	for i, x := range r {
		if i == 0 {
			last = x
			continue
		}
		if x < last {
			return false
		}
		last = x
	}
	return true
}

func (r reactorLevel) IsAdjacentSafe() bool {
	length := len(r)
	isOver := func(i int) bool {
		return i+1 > length-1
	}

	safe := func(t int) bool {
		return t >= 1 && t <= 3
	}
	testSafety := func(a int, testValue int) bool {
		t := a - testValue
		if t < 0 {
			t = -t
		}
		return safe(t)
	}

	for i, x := range r {
		if i == 0 {
			if isOver(i) {
				return true
			}
			if !testSafety(x, r[i+1]) {
				return false
			}
			continue
		} else if i == length-1 {
			break
		}

		if !testSafety(x, r[i-1]) {
			return false
		}
		if !testSafety(x, r[i+1]) {
			return false
		}
	}
	return true
}

func (r reactorLevel) IsSafe() bool {
	if !(r.IsReactorLevelDecreasing() || r.IsReactorLevelIncreasing()) {
		return false
	}
	return r.IsAdjacentSafe()
}

func CreateReactorLevel(input []string) reactorLevel {
	var ret reactorLevel
	for _, x := range input {
		val, err := strconv.Atoi(x)
		utils.Check(err)
		ret = append(ret, val)
	}
	return ret
}

func readInput(path string) []reactorLevel {
	f, err := os.Open(path)
	utils.Check(err)

	levels := []reactorLevel{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value := strings.Split(scanner.Text(), " ")
		levels = append(levels, CreateReactorLevel(value))
	}
	return levels
}

func Puzzle(path string) {
	levels := readInput(path)

	safe := 0
	for _, x := range levels {
		if x.IsSafe() {
			safe++
		}
	}

	fmt.Println("Puzzle 2 safety:", safe)
}

func permuateReactorLevels(r reactorLevel, permutation int) reactorLevel {
	newLevel := make(reactorLevel, len(r))
	copy(newLevel, r)
	newLevel = append(newLevel[:permutation], newLevel[permutation+1:]...)
	return newLevel
}

func PuzzleHard(path string) {
	levels := readInput(path)

	safe := 0

	for _, x := range levels {
		if x.IsSafe() {
			safe++
		} else {
			for i := range x {
				damped := permuateReactorLevels(x, i)
				if damped.IsSafe() {
					safe++
					break
				}
			}
		}
	}

	fmt.Println("Puzzle 2 safety with damping:", safe)
}
