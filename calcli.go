// A command line calculator program written in golang
package main

import (
	"flag"
	"fmt"
	"github.com/codythegreat/calcli/src"
	"math"
	"os"
	"strconv"
	"strings"
)

var userArgs string

// define flags
var abs = flag.Bool("abs", false, "converts return value to abs form")
var floor = flag.Bool("floor", false, "rounds result down")
var ceil = flag.Bool("ceil", false, "rounds result up")
var round = flag.Bool("round", false, "rounds result")
var latexI = flag.Bool("latexI", false, "Only prints LaTeX Inline formatting")
var latexD = flag.Bool("latexD", false, "Only prints LaTeX Display formatting")
var debug = flag.Bool("db", false, "prints functions, inputs, and outputs throughout execution")

func main() {
	// if os.Args missing equation, return error
	if !calclisrc.VerifyInputHasEquation(len(os.Args)) {
		fmt.Println("Please input an equation")
		os.Exit(1)
	}

	// strip os.Args for just equation
	userArgs = calclisrc.RemoveSpacesFromEquation(os.Args)

	// make sure that the equation uses proper syntax
	if !calclisrc.VerifyEquationHasProperSyntax(userArgs) {
		fmt.Println("Error: improper syntax found within equation")
		//TODO: print out a list of proper syntax
		os.Exit(1)
	}

	// if equation has improper curly brackets, break
	if !calclisrc.VerifyEquationHasProperBrackets(userArgs) {
		fmt.Println("Error: Missing curly brackets '{}' on one or more operators")
		os.Exit(1)
	}

	// removed equation from os.Args
	for i := 2; i < len(os.Args); i++ {
		os.Args[i] = ""
	}

	// check os.Args for flags, and set variables
	if len(os.Args) > 2 {
		flag.Parse()
	}

	// if latexI/D, simply print and quit
	// else print the result for given flag
	switch {
	case *latexI:
		fmt.Printf("$%s$\n", calclisrc.ConvertToLaTeX(userArgs))
	case *latexD:
		fmt.Printf("$$%s$$\n", calclisrc.ConvertToLaTeX(userArgs))
	case *floor:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Floor(solveEquationFloat(*debug)), 'f', -1, 64))
	case *ceil:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Ceil(solveEquationFloat(*debug)), 'f', -1, 64))
	case *round:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Round(solveEquationFloat(*debug)), 'f', -1, 64))
	case *abs:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Abs(solveEquationFloat(*debug)), 'f', -1, 64))
	default:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(solveEquationFloat(*debug), 'f', -1, 64))
	}
}

func printEquation() {
	// Print out equation:
	fmt.Printf("\n\n\tEquation:\t%s\n\n", userArgs)
	// Print out LaTeX translated equation
	fmt.Printf("\tLaTeX inline:\t$%s$\n\n", calclisrc.ConvertToLaTeX(userArgs))
	fmt.Printf("\tLateX Display:\t$$%s$$\n\n", calclisrc.ConvertToLaTeX(userArgs))
}

func solveEquationFloat(debug bool) float64 {
	floatAnswer, err := strconv.ParseFloat(calclisrc.ParseArgsParen(userArgs, debug), 64)
	if err != nil {
		fmt.Printf("While handling flags: %v", err)
	}
	return floatAnswer
}
