package main

import (
	"computorv1/src/polynomial"
	"computorv1/src/polynomialParser"
	"fmt"
	"os"
)

func main() {
    args := os.Args[1:]

	var polynomialInput string = args[0]
	pol := polynomialParser.Parse(polynomialInput)

	fmt.Println("Polynomial:", polynomial.Str(*pol))
	fmt.Println("Reduce each sides...")
	polynomial.Simplify(pol)
	fmt.Println("Reduced sides polynomial:", polynomial.Str(*pol))
	fmt.Println("Regroup all on left side...")
	polynomial.Regroup(pol)
	fmt.Println("Reduce...")
	polynomial.Simplify(pol)
	fmt.Println("Final reduced polynomial:", polynomial.Str(*pol))
}