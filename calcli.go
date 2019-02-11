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

func parseArgs(args string) string {
	// variables to hold logic of operations
	var hasExponent bool
	var hasMult bool
	var hasDiv bool
	var hasAdd bool
	var hasSub bool
	// regular expressions to interpret user input:
	floatRegex, err := regexp.Compile(`\d+\.\d+`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	intRegex, err := regexp.Compile(`\d+`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	plusOpRegex, err := regexp.Compile(`\+`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	subOpRegex, err := regexp.Compile(`\-`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	multOpRegex, err := regexp.Compile(`\*`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	divOpRegex, err := regexp.Compile(`\}`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	powerOpRegex, err := regexp.Compile(`\^`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	eNumbRegex, err := regexp.Compile(`e`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	piNumbRegex, err := regexp.Compile(`pi`)
	if err != nil {
		fmt.Printf("%v", err)
	}

	// using the regexps on the arguments
	floatIndex := floatRegex.FindAllIndex([]byte(args), -1)
	if floatIndex != nil {
		fmt.Println(args[floatIndex[0][0]:floatIndex[0][1]])
	}
	intIndex := intRegex.FindAllIndex([]byte(args), -1)
	if intIndex != nil {
		fmt.Println(args[intIndex[0][0]:intIndex[0][1]])
	}
	plusOpIndex := plusOpRegex.FindAllIndex([]byte(args), -1)
	if plusOpIndex != nil {
		fmt.Println(args[plusOpIndex[0][0]:plusOpIndex[0][1]])
		hasAdd = true
	}
	subOpIndex := subOpRegex.FindAllIndex([]byte(args), -1)
	if subOpIndex != nil {
		fmt.Println(args[subOpIndex[0][0]:subOpIndex[0][1]])
		hasSub = true
	}
	multOpIndex := multOpRegex.FindAllIndex([]byte(args), -1)
	if multOpIndex != nil {
		fmt.Println(args[multOpIndex[0][0]:multOpIndex[0][1]])
		hasMult = true
	}
	divOpIndex := divOpRegex.FindAllIndex([]byte(args), -1)
	if divOpIndex != nil {
		fmt.Println(args[divOpIndex[0][0]:divOpIndex[0][1]])
		hasDiv = true
	}
	powerOpIndex := powerOpRegex.FindAllIndex([]byte(args), -1)
	if powerOpIndex != nil {
		fmt.Println(args[powerOpIndex[0][0]:powerOpIndex[0][1]])
		hasExponent = true
	}
	eNumbIndex := eNumbRegex.FindAllIndex([]byte(args), -1)
	if eNumbIndex != nil {
		fmt.Println(args[eNumbIndex[0][0]:eNumbIndex[0][1]])
	}
	piNumbIndex := piNumbRegex.FindAllIndex([]byte(args), -1)
	if piNumbIndex != nil {
		fmt.Println(args[piNumbIndex[0][0]:piNumbIndex[0][1]])
	}

	if hasExponent == true {
		for _, power := range powerOpIndex {
			fString := floatRegex.FindString(args[:power[0]])
			d := intRegex.FindString(args[:power[0]])
			if fString != "" {
				f, err := strconv.ParseFloat(fString, 64)
				if err != nil {
					fmt.Printf("while parsing float in hasExponent: %v", err)
				}
				fpString := floatRegex.FindString(args[power[1]:])
				if fpString != "" {
					fp, err := strconv.ParseFloat(fpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasExponent: %v", err)
					}
					fmt.Printf("%f", math.Pow(f, fp))
				}
			} else if d != "" {
				fmt.Println("testing")
			}
		}
	}
	if hasExponent && hasMult && hasDiv && hasAdd && hasSub {
		fmt.Println("You have a lot of operators.")
	}
	fmt.Println(args)
	return args
}

func main() {
	// load arguments into variable
	parseArgsParen(userArgs, true)
	parseArgs(userArgs)
}
