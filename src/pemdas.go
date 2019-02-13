// handles the calculation of basic operations of the equation

package calclisrc

import (
	"fmt"
	"strconv"
	"strings"
)

func parseMult(loc []int, equation string) string {
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
	equation = equation[:loc[0]] + strconv.FormatFloat(leftSide*rightSide, 'f', -1, 64) + equation[loc[1]:]
	return equation
}
func parseDiv(loc []int, equation string) string {
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
	return equation
}
func parseAdd(loc []int, equation string) string {
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
	return equation
}
func parseSub(loc []int, equation string) string {
	// split at - to get both digits
	subString := strings.Split(equation[loc[0]:loc[1]], "-")
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
	return equation
}
