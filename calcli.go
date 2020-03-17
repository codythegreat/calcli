// A command line calculator program written in golang
package main

import (
	"flag"
	"fmt"
	"github.com/codythegreat/calcli/src"
	"math"
	"os"
	"regexp"
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

func main() {
	// if os.Args missing equation, return error
	if len(os.Args) == 1 {
		fmt.Println("Please input an equation")
		os.Exit(1)
	}
	// check os.Args for flags, and set variables
	flag.Parse()

	// strip os.Args for just equation
	if os.Args[1][:1] == "-" {
		userArgs = regexp.MustCompile(` `).ReplaceAllString(strings.Join(os.Args[2:len(os.Args)], ""), "")
	} else {
		userArgs = regexp.MustCompile(` `).ReplaceAllString(strings.Join(os.Args[1:len(os.Args)], ""), "")
	}

	// strip all spaces out of equation
	if os.Args[1][:1] == "-" {
		userArgs = regexp.MustCompile(` `).ReplaceAllString(strings.Join(os.Args[2:len(os.Args)], ""), "")
	} else {
		userArgs = regexp.MustCompile(` `).ReplaceAllString(strings.Join(os.Args[1:len(os.Args)], ""), "")
	}

	// if bad input, break
	if calclisrc.DetectInputError(userArgs) {
		fmt.Println("Error: Missing curly brackets '{}' on one or more operators")
		os.Exit(1)
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
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Floor(solveEquationFloat()), 'f', -1, 64))
	case *ceil:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Ceil(solveEquationFloat()), 'f', -1, 64))
	case *round:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Round(solveEquationFloat()), 'f', -1, 64))
	case *abs:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Abs(solveEquationFloat()), 'f', -1, 64))
	default:
		printEquation()
		userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(solveEquationFloat(), 'f', -1, 64))
	}
}

func printEquation() {
	// Print out equation:
	fmt.Printf("\n\n\tEquation:\t%s\n\n", userArgs)
	// Print out LaTeX translated equation
	fmt.Printf("\tLaTeX inline:\t$%s$\n\n", calclisrc.ConvertToLaTeX(userArgs))
	fmt.Printf("\tLateX Display:\t$$%s$$\n\n", calclisrc.ConvertToLaTeX(userArgs))
}

func solveEquationFloat() float64 {
	floatAnswer, err := strconv.ParseFloat(calclisrc.ParseArgsParen(userArgs), 64)
	if err != nil {
		fmt.Printf("While handling flags: %v", err)
	}
	return floatAnswer
}
