// handles the calculation of basic operations of the equation

package calclisrc

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseMult(loc []int, equation string, debug bool) string {
	if debug {
		DebugMessageFunctionStart(parseMult, equation)
	}
	// split at * to get both digits
	multString := strings.Split(equation[loc[0]:loc[1]], "*")
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(multString[0], 64)
	if err != nil {
		fmt.Printf("while parsing multiplication: %v", err)
	}
	rightSide, err := strconv.ParseFloat(multString[1], 64)
	if err != nil {
		fmt.Printf("while parsing multiplication: %v", err)
	}
	// write the values to the input
	if leftSide < 0 && rightSide < 0 {
		equation = equation[:loc[0]] + "-" + strconv.FormatFloat(math.Abs(leftSide*rightSide), 'f', -1, 64) + equation[loc[1]:]
		return equation
	}
	equation = equation[:loc[0]] + strconv.FormatFloat(leftSide*rightSide, 'f', -1, 64) + equation[loc[1]:]
	if debug {
		DebugMessageFunctionEnd(parseMult, equation)
	}
	return equation
}
func parseDiv(loc []int, equation string, debug bool) string {
	if debug {
		DebugMessageFunctionStart(parseDiv, equation)
	}
	// split at / to get both digits
	divString := strings.Split(equation[loc[0]:loc[1]], "/")
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(divString[0], 64)
	if err != nil {
		fmt.Printf("while parsing division: %v", err)
	}
	rightSide, err := strconv.ParseFloat(divString[1], 64)
	if err != nil {
		fmt.Printf("while parsing division: %v", err)
	}
	// write the values to the input
	equation = equation[:loc[0]] + strconv.FormatFloat(leftSide/rightSide, 'f', -1, 64) + equation[loc[1]:]
	if debug {
		DebugMessageFunctionEnd(parseDiv, equation)
	}
	return equation
}
func parseAdd(loc []int, equation string, debug bool) string {
	if debug {
		DebugMessageFunctionStart(parseAdd, equation)
	}
	// split at + to get both digits
	addString := strings.Split(equation[loc[0]:loc[1]], "+")
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(addString[0], 64)
	if err != nil {
		fmt.Printf("while parsing addition: %v", err)
	}
	rightSide, err := strconv.ParseFloat(addString[1], 64)
	if err != nil {
		fmt.Printf("while parsing addition: %v", err)
	}
	// write the values to the input
	equation = equation[:loc[0]] + strconv.FormatFloat(leftSide+rightSide, 'f', -1, 64) + equation[loc[1]:]
	if debug {
		DebugMessageFunctionEnd(parseAdd, equation)
	}
	return equation
}
func parseSub(loc []int, equation string, debug bool) string {
	if debug {
		DebugMessageFunctionStart(parseSub, equation)
	}
	// check to see if first number is negative. split at the minus sign.
	countMinus := strings.Count(equation, "-")
	var subString []string
	if countMinus > 1 {
		subString = strings.Split(equation[loc[0]+1:loc[1]], "-")
		subString[0] = "-" + subString[0]
	} else {
		subString = strings.Split(equation[loc[0]:loc[1]], "-")
	}
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(subString[0], 64)
	if err != nil {
		fmt.Printf("while parsing subtraction: %v", err)
	}
	rightSide, err := strconv.ParseFloat(subString[1], 64)
	if err != nil {
		fmt.Printf("while parsing subtraction: %v", err)
	}
	// write the values to the input
	equation = equation[:loc[0]] + strconv.FormatFloat(leftSide-rightSide, 'f', -1, 64) + equation[loc[1]:]
	if debug {
		DebugMessageFunctionEnd(parseSub, equation)
	}
	return equation
}
