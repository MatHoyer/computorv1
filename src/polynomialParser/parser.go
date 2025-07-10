package polynomialParser

import (
	"computorv1/src/number"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Parse(input string) {
	sides := strings.Split(input, "=")
	if len(sides) != 2 {
		panic("Wrong number of sides for the given polynomial")
	}

	for _, side := range(sides) {
		side = strings.ReplaceAll(side, " ", "")
		re := regexp.MustCompile(`[+\-]?\s*[^+\-]+`)
		parts := re.FindAllString(side, -1)
		for _, part := range(parts) {
			part = strings.ReplaceAll(part, "*", "")
			part = strings.ReplaceAll(part, "^", "")
			splitResult := strings.Split(part, "X")
			if len(splitResult) != 2 {
				panic(fmt.Sprintf("Invalid polynomial term: %s", part))
			}
			value, degree := splitResult[0], splitResult[1]

			f, err := strconv.ParseFloat(value, 32)
			if err != nil{
				panic(fmt.Sprintf("%s isn't a float", value))
			}

			i, err := strconv.Atoi(degree)
			if err != nil{
				panic(fmt.Sprintf("%s isn't a int", degree))
			}

			var n = number.Create(float32(f), i)
			fmt.Println(number.Str(*n))
		}
	}
}