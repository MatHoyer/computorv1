package polynomialParser

import (
	"computorv1/src/number"
	"fmt"
	"strconv"
	"strings"
)

func checkNumberOfRune(s string, ss string, count int) int {
	c := strings.Count(s, ss)
	if c > count {
		panic(fmt.Sprintf("too much %s in: %s", ss, s))
	}
	return c
}

func checkPrevRune(s string, ss string, pattern string) {
	i := strings.Index(s, ss)

	if i == 0 {
		panic(fmt.Sprintf("Bad token: %s", s))
	}
	v := string(s[i - 1])
	if !strings.Contains(pattern, v) {
		panic(fmt.Sprintf("Bad token: %s", s))
	}
}

func checkNextRune(s string, ss string, pattern string) {
	i := strings.Index(s, ss)

	if i == len(s) - 1 {
		panic(fmt.Sprintf("Bad token: %s", s))
	}
	v := string(s[i + 1])
	if !strings.Contains(pattern, v) {
		panic(fmt.Sprintf("Bad token: %s", s))
	}
}


func ParsePart(part string) *number.Number {
	var value, degree string

	cMult := checkNumberOfRune(part, "*", 1)
	cPow := checkNumberOfRune(part, "^", 1)
	cPoint := checkNumberOfRune(part, ".", 1)

	numbers := "0123456789"
	if cMult == 1 {
		checkPrevRune(part, "*", numbers)
		checkNextRune(part, "*", "X")
	}
	if cPow == 1 {
		checkPrevRune(part, "^", "X")
		checkNextRune(part, "^", numbers)
	}
	if cPoint == 1 {
		checkPrevRune(part, ".", numbers)
		checkNextRune(part, ".", numbers)
	}

	part = strings.ReplaceAll(part, "*", "")
	part = strings.ReplaceAll(part, "^", "")

	fmt.Println(part)
	if strings.Contains(part, "X") {
		splitted := strings.Split(part, "X")

		if len(splitted[0]) == 1 {
			switch splitted[0][0] {
			case '+':
				value = "+1"
			case '-':
				value = "-1"
			default:
				value = splitted[0]
			}
		} else if len(splitted[0]) == 0 {
			value = "+1"
		} else {
			value = splitted[0]
		}

		if len(splitted[1]) != 0 {
			degree = splitted[1]
		} else {
			degree = "1"
		}
	} else  {
		value = part
		degree = "0"
	}

	f, err := strconv.ParseFloat(value, 32)
	if err != nil{
		panic(fmt.Sprintf("%s isn't a float", value))
	}

	i, err := strconv.Atoi(degree)
	if err != nil{
		panic(fmt.Sprintf("%s isn't a int", degree))
	}

	return number.Create(float32(f), i)
}