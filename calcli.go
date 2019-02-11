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
	// variable to be returned
	returnString := args
	// variables to hold logic of operations
	var hasExponent bool
	var hasMult bool
	var hasDiv bool
	var hasAdd bool
	var hasSub bool
	// regular expressions to interpret user input:
	floatRegex, err := regexp.Compile(`^\d*\.\d+$`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	intRegex, err := regexp.Compile(`^\d+$`)
	if err != nil {
		fmt.Printf("%v", err)
	}
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
	// values used to write to return value
	var powerResults []float64
	if hasExponent == true { // if a '^' symbol exists in args
		for _, power := range powerOpIndex { // for each power found in FindAllIndex
			powerString := strings.Split(args[power[0]:power[1]], "^")
			fString := floatRegex.FindString(powerString[0])
			dString := intRegex.FindString(powerString[0])
			if fString != "" {
				f, err := strconv.ParseFloat(fString, 64)
				if err != nil {
					fmt.Printf("while parsing float in hasExponent: %v\n", err)
				}
				dpString := intRegex.FindString(powerString[1])
				fpString := floatRegex.FindString(powerString[1])
				if fpString != "" {
					fp, err := strconv.ParseFloat(fpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasExponent: %v\n", err)
					}
					fmt.Printf("%f\n", math.Pow(f, fp))
					powerResults = append(powerResults, math.Pow(f, fp))
				} else if dpString != "" {
					dp, err := strconv.ParseFloat(dpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasExponent: %v\n", err)
					}
					fmt.Printf("%f\n", math.Pow(f, dp))
					powerResults = append(powerResults, math.Pow(f, dp))
				}
			} else if dString != "" {
				d, err := strconv.ParseFloat(dString, 64)
				if err != nil {
					fmt.Printf("while parsing float in hasExponent: %v\n", err)
				}
				dpString := intRegex.FindString(powerString[1])
				fpString := floatRegex.FindString(powerString[1])
				if fpString != "" {
					fp, err := strconv.ParseFloat(fpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasExponent: %v\n", err)
					}
					fmt.Printf("%f\n", math.Pow(d, fp))
					powerResults = append(powerResults, math.Pow(d, fp))
				} else if dpString != "" {
					dp, err := strconv.ParseFloat(dpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasExponent: %v\n", err)
					}
					fmt.Printf("%f\n", math.Pow(d, dp))
					powerResults = append(powerResults, math.Pow(d, dp))
				}
			}
		}
		for i := len(powerResults) - 1; i >= 0; i-- {
			returnString = returnString[:powerOpIndex[i][0]] + strconv.FormatFloat(powerResults[i], 'f', -1, 64) + returnString[powerOpIndex[i][1]:]
		}
	}
	var multResults []float64
	if hasMult == true { // if a '*' symbol exists in args
		for _, multiple := range powerOpIndex { // for each power found in FindAllIndex
			multString := strings.Split(args[multiple[0]:multiple[1]], "^")
			fString := floatRegex.FindString(multString[0])
			dString := intRegex.FindString(multString[0])
			if fString != "" {
				f, err := strconv.ParseFloat(fString, 64)
				if err != nil {
					fmt.Printf("while parsing float in hasMult: %v\n", err)
				}
				dpString := intRegex.FindString(multString[1])
				fpString := floatRegex.FindString(multString[1])
				if fpString != "" {
					fp, err := strconv.ParseFloat(fpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasMult: %v\n", err)
					}
					fmt.Printf("%f\n", math.Pow(f, fp))
					multResults = append(multResults, f*fp)
				} else if dpString != "" {
					dp, err := strconv.ParseFloat(dpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasMult: %v\n", err)
					}
					fmt.Printf("%f\n", math.Pow(f, dp))
					multResults = append(multResults, f*dp)
				}
			} else if dString != "" {
				d, err := strconv.ParseFloat(dString, 64)
				if err != nil {
					fmt.Printf("while parsing float in hasMult: %v\n", err)
				}
				dpString := intRegex.FindString(multString[1])
				fpString := floatRegex.FindString(multString[1])
				if fpString != "" {
					fp, err := strconv.ParseFloat(fpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasMult: %v\n", err)
					}
					fmt.Printf("%f\n", d*fp)
					multResults = append(multResults, d*fp)
				} else if dpString != "" {
					dp, err := strconv.ParseFloat(dpString, 64)
					if err != nil {
						fmt.Printf("while parsing float in hasMult: %v\n", err)
					}
					fmt.Printf("%f\n", d*dp)
					multResults = append(multResults, d*dp)
				}
			}
		}
		for i := len(multResults) - 1; i >= 0; i-- {
			returnString = returnString[:multOpIndex[i][0]] + strconv.FormatFloat(multResults[i], 'f', -1, 64) + returnString[multOpIndex[i][1]:]
		}
	}
	if hasExponent && hasMult && hasDiv && hasAdd && hasSub {
		fmt.Println("You have a lot of operators.")
	}
	fmt.Println(returnString)
	return returnString
}

func main() {
	// load arguments into variable
	parseArgsParen(userArgs, true)
	parseArgs(userArgs)
}
