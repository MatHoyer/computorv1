package solver

import (
	"computorv1/src/expression"
	"computorv1/src/number"
	"computorv1/src/polynomial"
	"fmt"
	"math"
)

func Solve(pol polynomial.Polynomial) {
	switch expression.GetHightestDegree(*pol.Left) {
	case 0:
		if expression.Str(*pol.Left) == expression.Str(*pol.Right) {
			fmt.Println("Any real number is a solution")
		} else {
			fmt.Println("No solution")
		}
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

	a, ok := left.Values[2]
	if !ok {
		a = number.Create(0, 2)
	}
	b, ok := left.Values[1]
	if !ok {
		b = number.Create(0, 1)
	}
	c, ok := left.Values[0]
	if !ok {
		c = number.Create(0, 0)
	}

	fmt.Println("  •", fmt.Sprintf(
		"%g^2 - 4 * %g * %g",
		b.Value,
		c.Value,
		a.Value,
	))
	delta := math.Pow(float64(b.Value), 2)
	delta -= 4 * float64(c.Value) * float64(a.Value) 
	fmt.Println("  •", fmt.Sprintf("%g", delta))

	if delta > 0 {
		fmt.Println("Delta is strictly positive, so there are 2 real solutions:")
	
		sol1 := (-b.Value - float32(math.Sqrt(delta))) / (2 * a.Value)
		sol2 := (-b.Value + float32(math.Sqrt(delta))) / (2 * a.Value)
		fmt.Printf("%g\n", sol1)
		fmt.Printf("%g\n", sol2)
	} else if delta == 0 {
		fmt.Println("Delta is equal to 0, so there is 1 real solution:")
	
		sol := -b.Value / (2 * a.Value)
		fmt.Printf("%g\n", sol)
	} else {
		fmt.Println("Delta is strictly negative, so there are 2 complex solutions:")
		
		delta = -delta
		denominator := (2 * a.Value)
		ratio := int(b.Value) % int(denominator)

		firstNumerator := (-b.Value) / float32(ratio)
		secondNumerator := float32(math.Sqrt(delta)) / float32(ratio)
		denominator = denominator / float32(ratio)

		fmt.Printf("%g/%g - %gi/%g\n", firstNumerator, denominator, secondNumerator, denominator)
		fmt.Printf("%g/%g + %gi/%g\n", firstNumerator, denominator, secondNumerator, denominator)
	}
}