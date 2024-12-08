package puzzle3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tfx-aoc24/utils"
	"unicode"
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

type expectedChar int

const (
	M_CHAR expectedChar = iota
	U_CHAR
	L_CHAR
	START_ARGS
	ARGS_1
	ARGS_2
)

type expectedConditional int

const (
	D_ECD expectedConditional = iota
	O_ECD
	N_ECD
	APOS_ECD
	T_ECD
	START_ARGS_ECD
	END_ARGS_ECD
)

func IsExpectedChar(input rune, expected expectedChar) bool {
	switch expected {
	case M_CHAR:
		return input == 'm'
	case U_CHAR:
		return input == 'u'
	case L_CHAR:
		return input == 'l'
	case START_ARGS:
		return input == '('
	case ARGS_1:
		return unicode.IsDigit(input) || input == ','
	case ARGS_2:
		return unicode.IsDigit(input) || input == ')'
	}
	return false
}

func IsExpectedConditional(input rune, expected expectedConditional) bool {
	switch expected {
	case D_ECD:
		return input == 'd'
	case O_ECD:
		return input == 'o'
	case N_ECD:
		return input == 'n' || input == '('
	case APOS_ECD:
		return input == '\''
	case T_ECD:
		return input == 't'
	case START_ARGS_ECD:
		return input == '('
	case END_ARGS_ECD:
		return input == ')'
	}
	return false
}

type mulInstruction struct {
	Instruction  string
	A            string
	B            string
	Expected     expectedChar
	Error        bool
	Constructed  bool
	ArgCompleted bool
}

func (m mulInstruction) Value() int {
	a, err := strconv.Atoi(m.A)
	utils.Check(err)
	b, err := strconv.Atoi(m.B)
	utils.Check(err)
	return a * b
}

func (m *mulInstruction) ReadNext(input rune) {
	if !IsExpectedChar(input, m.Expected) {
		m.Error = true
		return
	}

	m.Instruction += string(input)

	switch m.Expected {
	case M_CHAR:
		m.Expected = U_CHAR
	case U_CHAR:
		m.Expected = L_CHAR
	case L_CHAR:
		m.Expected = START_ARGS
	case START_ARGS:
		m.Expected = ARGS_1
	case ARGS_1:
		if input == ',' {
			m.Expected = ARGS_2
			m.ArgCompleted = false
			return
		}
		m.A += string(input)
		m.ArgCompleted = true
	case ARGS_2:
		if input == ')' && m.ArgCompleted {
			m.Constructed = true
			return
		} else if !unicode.IsDigit(input) {
			fmt.Println("failed due to digit error")
			m.Error = true
			return
		}
		m.B += string(input)
		m.ArgCompleted = true
	}
}

type conditionalInstruction struct {
	Instruction string
	Expected    expectedConditional
	Error       bool
	Completed   bool
}

func (c *conditionalInstruction) ReadNext(input rune) {
	if !IsExpectedConditional(input, c.Expected) {
		c.Error = true
		return
	}

	c.Instruction += string(input)

	switch c.Expected {
	case D_ECD:
		c.Expected = O_ECD
	case O_ECD:
		c.Expected = N_ECD
	case N_ECD:
		if input == '(' {
			c.Expected = END_ARGS_ECD
		} else {
			c.Expected = APOS_ECD
		}
	case APOS_ECD:
		c.Expected = T_ECD
	case T_ECD:
		c.Expected = START_ARGS_ECD
	case START_ARGS_ECD:
		c.Expected = END_ARGS_ECD
	case END_ARGS_ECD:
		c.Completed = true
	}
}

func (c *conditionalInstruction) Value() bool {
	return c.Instruction != "do()"
}

func scanLine(input string, ignore bool) (int, bool) {
	currInst := mulInstruction{}
	currConditional := conditionalInstruction{}
	useConditional := false
	sum := 0
	ignoreValue := ignore
	for _, x := range input {
		if !useConditional {
			currInst.ReadNext(x)
			if currInst.Error {
				// reset instruction
				currInst = mulInstruction{}
				if x == 'd' {
					useConditional = true
					currConditional = conditionalInstruction{}
					currConditional.ReadNext(x)
				}
			} else if currInst.Constructed {
				if !ignoreValue {
					fmt.Println("value read:", currInst.Value(), currInst)
					sum += currInst.Value()
				}
				currInst = mulInstruction{}
			}
			continue
		}

		currConditional.ReadNext(x)
		if currConditional.Error {
			useConditional = false
			currConditional = conditionalInstruction{}
			if x == 'm' {
				currInst = mulInstruction{}
				currInst.ReadNext(x)
			}
			continue
		} else if currConditional.Completed {
			ignoreValue = currConditional.Value()
			fmt.Println("Conditional complete: now ignoring value?", ignoreValue)
			currConditional = conditionalInstruction{}
		}
	}
	return sum, ignoreValue
}

// answer is not: 879544237516
func Puzzle(path string) {
	input := readInput(path)
	sum := 0
	ignore := false
	for _, x := range input {
		val, ign := scanLine(x, ignore)
		ignore = ign
		sum += val
	}

	fmt.Println("Puzzle 3 sum:", sum)
}
