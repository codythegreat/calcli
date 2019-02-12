// A command line calculator program written in golang
package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// grab equation and strip all spaces out of it
var userArgs = regexp.MustCompile(` `).ReplaceAllString(strings.Join(os.Args[len(os.Args)-1:], ""), "")

// define flags
var abs = flag.Bool("abs", false, "converts return value to abs form")
var floor = flag.Bool("floor", false, "rounds result down")
var ceil = flag.Bool("ceil", false, "rounds result up")
var round = flag.Bool("round", false, "rounds result")

func parseArgsParen(args string) string {
	// defile parentheses regex (only finds inner parentheses)
	parenOpRegex, err := regexp.Compile(`\([^\(\)]+\)`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// parse all parentheses, starting from innermost and working up to outermost
	for parenIndexes := parenOpRegex.FindStringIndex(args); parenIndexes != nil; parenIndexes = parenOpRegex.FindStringIndex(args) {
		args = args[:parenIndexes[0]] + parseArgs(args[parenIndexes[0]+1:parenIndexes[1]-1]) + args[parenIndexes[1]:]
	}
	// Once we know that all parentheses are resolved, parse other args and return string answer
	return parseArgs(args)
}

func parseSqrt(loc []int, equation string) string {
	// strip SQRT declaration from equation
	innerSqrt := equation[loc[0]+5 : loc[1]-1]
	// if SQRT[] contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, innerSqrt)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		innerSqrt = parseArgsParen(innerSqrt)
	}
	// if SQRT[] only contains a number, parse it as float
	sqrtFloat, err := strconv.ParseFloat(innerSqrt, 64)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// parse both digits in equation
	equation = equation[:loc[0]] + strconv.FormatFloat(math.Sqrt(sqrtFloat), 'f', -1, 64) + equation[loc[1]:]
	return equation
}

func parsePower(loc []int, equation string) string {
	// get right side of power
	rightSide := equation[strings.Index(equation, "{")+1 : loc[1]-1]
	leftSide := equation[loc[0]:strings.Index(equation, "^")]
	// if power contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, rightSide)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		rightSide = parseArgsParen(rightSide)
	}
	// parse rightSide into float
	rightSideFloat, err := strconv.ParseFloat(rightSide, 64)
	if err != nil {
		fmt.Printf("%v", err)
	}
	leftSideFloat, err := strconv.ParseFloat(leftSide, 64)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// parse both digits in equation
	equation = equation[:loc[0]] + strconv.FormatFloat(math.Pow(leftSideFloat, rightSideFloat), 'f', -1, 64) + equation[loc[1]:]
	return equation
}
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

func parseArgs(args string) string {
	// variable to be returned
	returnString := args

	// regular expressions to interpret user input:
	addOpRegex, err := regexp.Compile(`\d+(\.\d*)?\+\d+(\.\d*)?`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	subOpRegex, err := regexp.Compile(`\d+(\.\d*)?\-\d+(\.\d*)?`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	multOpRegex, err := regexp.Compile(`\d+(\.\d*)?\*\d+(\.\d*)?`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	divOpRegex, err := regexp.Compile(`\d+(\.\d*)?/\d+(\.\d*)?`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// starts from the innermost power in the equation
	powerOpRegex, err := regexp.Compile(`\d+(\.\d*)?\^\{[^\{\}]*\}`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// starts from the innermost sqrt in the equation
	sqrtOpRegex, err := regexp.Compile(`SQRT\{[^\{\}]*\}`)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for b := strings.Contains(returnString, "{"); b != false; b = strings.Contains(returnString, "{") {
		// parse all sqrt in equation
		for sqrtOpIndex := sqrtOpRegex.FindStringIndex(returnString); sqrtOpIndex != nil; sqrtOpIndex = sqrtOpRegex.FindStringIndex(returnString) {
			returnString = parseSqrt(sqrtOpIndex, returnString)
		}
		// parse all exponents in equation
		for powerOpIndex := powerOpRegex.FindStringIndex(returnString); powerOpIndex != nil; powerOpIndex = powerOpRegex.FindStringIndex(returnString) {
			returnString = parsePower(powerOpIndex, returnString)
		}
	}
	// parse all multiplication in equation
	for multOpIndex := multOpRegex.FindStringIndex(returnString); multOpIndex != nil; multOpIndex = multOpRegex.FindStringIndex(returnString) {
		returnString = parseMult(multOpIndex, returnString)
	} // parse all division in equation
	for divOpIndex := divOpRegex.FindStringIndex(returnString); divOpIndex != nil; divOpIndex = divOpRegex.FindStringIndex(returnString) {
		returnString = parseDiv(divOpIndex, returnString)
	}
	// parse all addition in equation
	for addOpIndex := addOpRegex.FindStringIndex(returnString); addOpIndex != nil; addOpIndex = addOpRegex.FindStringIndex(returnString) {
		returnString = parseAdd(addOpIndex, returnString)
	}
	// parse all subtraction in equation
	for subOpIndex := subOpRegex.FindStringIndex(returnString); subOpIndex != nil; subOpIndex = subOpRegex.FindStringIndex(returnString) {
		returnString = parseSub(subOpIndex, returnString)
	}

	// return string after parsing through order of operations
	return returnString
}

func convertToLaTeX(equation string) string {
	sqrtRegex, err := regexp.Compile(`SQRT`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	equation = sqrtRegex.ReplaceAllString(equation, "\\sqrt")
	return equation
}

func main() {
	// check os.Args for flags, and set variables
	flag.Parse()
	// Print out equation:
	fmt.Printf("Begining equation: %s\n", userArgs)
	fmt.Printf("LaTeX inline: $%s$\n", convertToLaTeX(userArgs))
	fmt.Printf("Latex Display: $$%s$$\n", convertToLaTeX(userArgs))
	// Print out answer:
	floatAnswer, err := strconv.ParseFloat(parseArgsParen(userArgs), 64)
	if err != nil {
		fmt.Printf("While handling flags: %v", err)
	}
	switch {
	case *floor:
		fmt.Printf("return value: %v\n", strconv.FormatFloat(math.Floor(floatAnswer), 'f', -1, 64))
	case *ceil:
		fmt.Printf("return value: %v\n", strconv.FormatFloat(math.Ceil(floatAnswer), 'f', -1, 64))
	case *round:
		fmt.Printf("return value: %v\n", strconv.FormatFloat(math.Round(floatAnswer), 'f', -1, 64))
	case *abs:
		fmt.Printf("return value: %v\n", strconv.FormatFloat(math.Abs(floatAnswer), 'f', -1, 64))
	default:
		fmt.Printf("return value: %v\n", parseArgsParen(userArgs))
	}
}
