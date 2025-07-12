package polynomialParser

import (
	"computorv1/src/number"
	"fmt"
	"strconv"
	"strings"
)

func ParsePart(part string) *number.Number {
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

	return number.Create(float32(f), i)
}