package main

import (
	"computorv1/src/expression"
	"computorv1/src/polynomial"
	"computorv1/src/polynomialParser"
	"computorv1/src/solver"
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}()

    args := os.Args[1:]

	var polynomialInput string = args[0]
	pol := polynomialParser.Parse(polynomialInput)

	fmt.Println("Polynomial:", polynomial.Str(*pol))

	fmt.Println("\nReduce each sides...")
	polynomial.Simplify(pol)
	fmt.Println("Reduced sides polynomial:", polynomial.Str(*pol))

	fmt.Println("\nRegroup all on left side...")
	polynomial.Regroup(pol)

	fmt.Println("\nFinal reduce...")
	polynomial.Simplify(pol)
	fmt.Println("Final reduced polynomial:", polynomial.Str(*pol))
	fmt.Println("Polynomial degree:", expression.GetHightestDegree(*pol.Left))

	fmt.Println("\nSolve...")
	solver.Solve(*pol)

}