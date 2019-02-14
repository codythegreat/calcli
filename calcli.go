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

// grab equation and strip all spaces out of it
var userArgs = regexp.MustCompile(` `).ReplaceAllString(strings.Join(os.Args[len(os.Args)-1:], ""), "")

// define flags
var abs = flag.Bool("abs", false, "converts return value to abs form")
var floor = flag.Bool("floor", false, "rounds result down")
var ceil = flag.Bool("ceil", false, "rounds result up")
var round = flag.Bool("round", false, "rounds result")

func main() {
	// check os.Args for flags, and set variables
	flag.Parse()
	// Print out equation:
	fmt.Printf("\n\n\tEquation:\t%s\n\n", userArgs)
	// Print out LaTeX translated equation
	fmt.Printf("\tLaTeX inline:\t$%s$\n\n", calclisrc.ConvertToLaTeX(userArgs))
	fmt.Printf("\tLateX Display:\t$$%s$$\n\n", calclisrc.ConvertToLaTeX(userArgs))
	// add parentheses so that innermost values are calculated first
	userArgs = strings.Replace(strings.Replace(userArgs, "{", "{(", -1), "}", ")}", -1)
	// Print out answer:
	floatAnswer, err := strconv.ParseFloat(calclisrc.ParseArgsParen(userArgs), 64)
	if err != nil {
		fmt.Printf("While handling flags: %v", err)
	}
	switch {
	case *floor:
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Floor(floatAnswer), 'f', -1, 64))
	case *ceil:
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Ceil(floatAnswer), 'f', -1, 64))
	case *round:
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Round(floatAnswer), 'f', -1, 64))
	case *abs:
		fmt.Printf("\treturn value:\t%v\n\n", strconv.FormatFloat(math.Abs(floatAnswer), 'f', -1, 64))
	default:
		fmt.Printf("\treturn value:\t%v\n\n", calclisrc.ParseArgsParen(userArgs))
	}
}
