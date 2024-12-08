package puzzle4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tfx-aoc24/utils"
)

func readInput(path string) []string {
	f, err := os.Open(path)
	utils.Check(err)

	ret := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

type direction int

const (
	NORTH direction = iota
	SOUTH
	WEST
	EAST
	NORTH_WEST
	NORTH_EAST
	SOUTH_WEST
	SOUTH_EAST
)

type match struct {
	pairs [][]int
	dir   direction
}

func debugGetDirName(dir direction) string {
	switch dir {
	case NORTH:
		return "NORTH"
	case SOUTH:
		return "SOUTH"
	case WEST:
		return "WEST"
	case EAST:
		return "EAST"
	case NORTH_WEST:
		return "NORTH_WEST"
	case NORTH_EAST:
		return "NORTH_EAST"
	case SOUTH_WEST:
		return "SOUTH_WEST"
	case SOUTH_EAST:
		return "SOUTH_EAST"
	}
	return "INVALID DIRECTION"
}

func getDirectionCoords(dir direction) (int, int) {
	switch dir {
	case NORTH:
		return 0, -1
	case SOUTH:
		return 0, 1
	case WEST:
		return -1, 0
	case EAST:
		return 1, 0
	case NORTH_WEST:
		return -1, -1
	case NORTH_EAST:
		return 1, -1
	case SOUTH_WEST:
		return -1, 1
	case SOUTH_EAST:
		return 1, 1
	}
	return 0, 0
}

func testDirection(input [][]rune, x, y, lenX, lenY int, expectedChar rune) bool {
	if x < 0 || x >= lenX-1 {
		return false
	}
	if y < 0 || y > lenY-1 {
		return false
	}
	return input[y][x] == expectedChar
}

var xmasString = [4]rune{'X', 'M', 'A', 'S'}

const targetStringInd = 3

var pairsOut = [][][]int{}
var matches = []match{}

func parseMapFromPosition(input [][]rune, x, y int) int {
	// this might be messed up...
	hits := 0
	lenY := len(input)
	for dir := NORTH; dir <= SOUTH_EAST; dir++ {
		lenX := len(input[x])
		dx, dy := getDirectionCoords(dir)
		nx := x
		ny := y

		xmasIndex := 0
		pairs := [][]int{}
		for testDirection(input, nx, ny, lenX, lenY, xmasString[xmasIndex]) {
			pairs = append(pairs, []int{nx, ny})
			xmasIndex++
			nx += dx
			ny += dy
			if xmasIndex > targetStringInd {
				// woo! we got it!
				fmt.Println("XMAS found as", debugGetDirName(dir), pairs)
				pairsOut = append(pairsOut, pairs)
				matches = append(matches, match{pairs: pairs, dir: dir})
				hits++
				break
			}
		}
	}
	return hits
}

// 2820 - too high
// 2626 - too high
// 2634 - too high...

// 2579 - too low!

func Puzzle(path string) {
	input := readInput(path)
	thing := [][]rune{}
	for _, x := range input {
		thing = append(thing, []rune(x))
	}

	hits := 0
	for y, row := range thing {
		for x, _ := range row {
			hits += parseMapFromPosition(thing, x, y)
		}
	}

	f, err := os.Create("p4output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	write := bufio.NewWriter(f)
	for _, p := range matches {
		strs := []string{}
		for _, pairs := range p.pairs {
			strs = append(strs, fmt.Sprintf("%d,%d", pairs[0], pairs[1]))
		}
		write.WriteString(debugGetDirName(p.dir) + "|" + strings.Join(strs, ";") + "\n")
		write.Flush()
	}

	fmt.Println("Puzzle 4 hits: ", hits)
}
