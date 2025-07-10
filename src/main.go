package main

import (
	"computorv1/src/polynomialParser"
	"os"
)

func main() {
    args := os.Args[1:]

	var polynomial string = args[0]
	polynomialParser.Parse(polynomial)
}