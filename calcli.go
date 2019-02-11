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

var userArgs = strings.Join(os.Args[1:], "")

var floor = flag.Bool("floor", false, "rounds result down")
var ceil = flag.Bool("ceil", false, "rounds result up")
var round = flag.Bool("round", false, "rounds result")

func parseArgsParen(args string, fullString bool) string { // fullString should be set to true when calling with all arguments, otherwise false
	parenOpRegex, err := regexp.Compile(`\(.+\)`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	parenIndexes := parenOpRegex.FindAllIndex([]byte(args), -1)
	if parenIndexes != nil {
		for _, match := range parenIndexes {
			fmt.Println(args[match[0]:match[1]])
			args = args[:match[0]] + string(parseArgsParen(args[match[0]+1:match[1]-1], false)) + args[:match[1]]
		}
	}
	fmt.Printf("Arguments: %s\n", args)
	return parseArgs(args)
}

func parsePower(loc []int, equation string) string {
	// split at ^ go get both digits
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

func parseArgs(args string) string {
	// variable to be returned
	returnString := args
	// regular expressions to interpret user input:
	plusOpRegex, err := regexp.Compile(`\d+(\.\d*)?\+\d+(\.\d*)?`)
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

	// using the regexps on the arguments
	plusOpIndex := plusOpRegex.FindStringIndex(returnString)
	if plusOpIndex != nil {
		fmt.Println(args[plusOpIndex[0]:plusOpIndex[1]])
	}
	subOpIndex := subOpRegex.FindStringIndex(returnString)
	if subOpIndex != nil {
		fmt.Println(args[subOpIndex[0]:subOpIndex[1]])
	}
	divOpIndex := divOpRegex.FindStringIndex(returnString)
	if divOpIndex != nil {
		fmt.Println(args[divOpIndex[0]:divOpIndex[1]])
	}
	powerOpIndex := powerOpRegex.FindStringIndex(returnString)
	if powerOpIndex != nil {
		fmt.Println(args[powerOpIndex[0]:powerOpIndex[1]])
	}
	// values used to write to return value
	for powerOpIndex := powerOpRegex.FindStringIndex(returnString); powerOpIndex != nil; powerOpIndex = powerOpRegex.FindStringIndex(returnString) {
		returnString = parsePower(powerOpIndex, returnString)
	}
	multOpIndex := multOpRegex.FindAllIndex([]byte(returnString), -1)
	if multOpIndex != nil {
		fmt.Println(returnString[multOpIndex[0][0]:multOpIndex[0][1]])
	}
	fmt.Println(returnString)
	return returnString
}

func main() {
	// load arguments into variable
	parseArgsParen(userArgs, true)
}
