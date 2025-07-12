package polynomialParser

import (
	"regexp"
	"strings"
)

func ParseSide(side string) []string {
	side = strings.ReplaceAll(side, " ", "")
	re := regexp.MustCompile(`[+\-]?\s*[^+\-]+`)
	parts := re.FindAllString(side, -1)

	return parts
}