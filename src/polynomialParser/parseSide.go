package polynomialParser

import (
	"regexp"
)

func ParseSide(side string) []string {
	re := regexp.MustCompile(`[+\-]?\s*[^+\-]+`)
	parts := re.FindAllString(side, -1)

	return parts
}