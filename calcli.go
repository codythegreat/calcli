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
	parenOpRegex, err := regexp.Compile(`\(\-?\d+(\.\d+)?((\^|\*|\/|\+|\-)\-?\d+(\.\d+)?)+\)`)
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

func parsePower(loc []int, equation string) string {
	// split at ^ to get both digits
	powerString := strings.Split(equation[loc[0]:loc[1]], "^")
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(powerString[0], 64)
	if err != nil {
		fmt.Printf("while parsing exponents: %v", err)
	}
	rightSide, err := strconv.ParseFloat(powerString[1], 64)
	if err != nil {
		fmt.Printf("while parsing exponents: %v", err)
	}
	// write the values to the input
	equation = equation[:loc[0]] + strconv.FormatFloat(math.Pow(leftSide, rightSide), 'f', -1, 64) + equation[loc[1]:]
	return equation
}
func parseMult(loc []int, equation string) string {
	// split at * to get both digits
	powerString := strings.Split(equation[loc[0]:loc[1]], "*")
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(powerString[0], 64)
	if err != nil {
		fmt.Printf("while parsing multiplication: %v", err)
	}
	rightSide, err := strconv.ParseFloat(powerString[1], 64)
	if err != nil {
		fmt.Printf("while parsing multiplication: %v", err)
	}
	// write the values to the input
	equation = equation[:loc[0]] + strconv.FormatFloat(leftSide*rightSide, 'f', -1, 64) + equation[loc[1]:]
	return equation
}
func parseDiv(loc []int, equation string) string {
	// split at / to get both digits
	powerString := strings.Split(equation[loc[0]:loc[1]], "/")
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(powerString[0], 64)
	if err != nil {
		fmt.Printf("while parsing division: %v", err)
	}
	rightSide, err := strconv.ParseFloat(powerString[1], 64)
	if err != nil {
		fmt.Printf("while parsing division: %v", err)
	}
	// write the values to the input
	equation = equation[:loc[0]] + strconv.FormatFloat(leftSide/rightSide, 'f', -1, 64) + equation[loc[1]:]
	return equation
}
func parseAdd(loc []int, equation string) string {
	// split at + to get both digits
	powerString := strings.Split(equation[loc[0]:loc[1]], "+")
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(powerString[0], 64)
	if err != nil {
		fmt.Printf("while parsing addition: %v", err)
	}
	rightSide, err := strconv.ParseFloat(powerString[1], 64)
	if err != nil {
		fmt.Printf("while parsing addition: %v", err)
	}
	// write the values to the input
	equation = equation[:loc[0]] + strconv.FormatFloat(leftSide+rightSide, 'f', -1, 64) + equation[loc[1]:]
	return equation
}
func parseSub(loc []int, equation string) string {
	// split at - to get both digits
	powerString := strings.Split(equation[loc[0]:loc[1]], "-")
	// parse both digits in equation
	leftSide, err := strconv.ParseFloat(powerString[0], 64)
	if err != nil {
		fmt.Printf("while parsing subtraction: %v", err)
	}
	rightSide, err := strconv.ParseFloat(powerString[1], 64)
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
	powerOpRegex, err := regexp.Compile(`\d+(\.\d*)?\^\d+(\.\d*)?`)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// parse all exponents in equation
	for powerOpIndex := powerOpRegex.FindStringIndex(returnString); powerOpIndex != nil; powerOpIndex = powerOpRegex.FindStringIndex(returnString) {
		returnString = parsePower(powerOpIndex, returnString)
	}
	// parse all multiplication in equation
	for multOpIndex := multOpRegex.FindStringIndex(returnString); multOpIndex != nil; multOpIndex = multOpRegex.FindStringIndex(returnString) {
		returnString = parseMult(multOpIndex, returnString)
	}
	// parse all division in equation
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

func main() {
	// check os.Args for flags, and set variables
	flag.Parse()
	// Print out equation:
	fmt.Printf("Begining equation: %s\n", userArgs)
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
