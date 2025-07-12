package polynomialParser

import (
	"computorv1/src/expression"
	"computorv1/src/polynomial"
	"strings"
)

func Parse(input string) *polynomial.Polynomial {
	sides := strings.Split(input, "=")
	if len(sides) != 2 {
		panic("Wrong number of sides for the given polynomial")
	}

	expL := expression.Create("+")
	expR := expression.Create("+")
	pol := polynomial.Create(expL, expR)

	var leftParts []string = ParseSide(sides[0])
	for _, part := range(leftParts) {
		expression.Append(expL, ParsePart(part))
	}

	var rightParts []string = ParseSide(sides[1])
	for _, part := range(rightParts) {
		expression.Append(expR, ParsePart(part))
	}

	return pol
}