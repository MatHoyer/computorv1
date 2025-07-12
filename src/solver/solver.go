package solver

import (
	"computorv1/src/expression"
	"computorv1/src/polynomial"
	"fmt"
	"math"
)

func Solve(pol polynomial.Polynomial) {
	switch expression.GetHightestDegree(*pol.Left) {
	case 0:
		fmt.Println("No solution")
	case 1:
		solveDegree1(pol)
	case 2:
		solveDegree2(pol)
	default:
		fmt.Println("The polynomial degree is strictly greater than 2, I can't solve")
	}
}

func solveDegree1(pol polynomial.Polynomial) {
	polynomial.BasicTransfer(
		pol.Left,
		pol.Right,
		0,
	)
	fmt.Println(" •", polynomial.Str(pol))

	pol.Right.Operation = "/"
	polynomial.DivideTransfer(pol.Left, pol.Right, 1)
	fmt.Println(" •", polynomial.Str(pol))
	fmt.Println("Solution is exactly:", polynomial.Str(pol))
	expression.Simplify(pol.Right)
	fmt.Println("and nearly:", polynomial.Str(pol))
}

func solveDegree2(pol polynomial.Polynomial) {
	left := pol.Left
	fmt.Println(" Delta:")
	fmt.Println("  •", fmt.Sprintf(
		"%g^2 - 4 * %g * %g",
		left.Values[1].Value,
		left.Values[0].Value,
		left.Values[2].Value,
	))
	delta := math.Pow(float64(left.Values[1].Value), 2)
	delta -= 4 * float64(left.Values[0].Value) * float64(left.Values[2].Value) 
	fmt.Println("  •", fmt.Sprintf("%g", delta))

	if delta > 0 {
		fmt.Println("Delta is strictly positive, so there are 2 real solutions")
	} else if delta == 0 {
		fmt.Println("Delta is equal to 0, so there is 1 real solution")
	} else {
		fmt.Println("Delta is strictly negative, so there are 2 complex solutions")
	}
}