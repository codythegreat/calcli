// recieves a function operator from parse.go, resolves it to a single float value, returns it
package calclisrc

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func parseSin(loc []int, equation string, debug bool) string {
	// strip sqrt declaration from equation
	innerSin := equation[loc[0]+4 : loc[1]-1]
	// if sqrt[] contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, innerSin)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		innerSin = ParseArgsParen(innerSin, debug)
	}
	// if sqrt[] only contains a number, parse it as float
	sinFloat, err := strconv.ParseFloat(innerSin, 64)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// parse both digits in equation
	equation = equation[:loc[0]] + strconv.FormatFloat(math.Sin(sinFloat), 'f', -1, 64) + equation[loc[1]:]
	return equation
}
func parseTan(loc []int, equation string, debug bool) string {
	// strip sqrt declaration from equation
	innerTan := equation[loc[0]+4 : loc[1]-1]
	// if sqrt[] contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, innerTan)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		innerTan = ParseArgsParen(innerTan, debug)
	}
	// if sqrt[] only contains a number, parse it as float
	tanFloat, err := strconv.ParseFloat(innerTan, 64)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// parse both digits in equation
	equation = equation[:loc[0]] + strconv.FormatFloat(math.Tan(tanFloat), 'f', -1, 64) + equation[loc[1]:]
	return equation
}
func parseCos(loc []int, equation string, debug bool) string {
	// strip sqrt declaration from equation
	innerCos := equation[loc[0]+4 : loc[1]-1]
	// if sqrt[] contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, innerCos)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		innerCos = ParseArgsParen(innerCos, debug)
	}
	// if sqrt[] only contains a number, parse it as float
	cosFloat, err := strconv.ParseFloat(innerCos, 64)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// parse both digits in equation
	equation = equation[:loc[0]] + strconv.FormatFloat(math.Cos(cosFloat), 'f', -1, 64) + equation[loc[1]:]
	return equation
}
func parsePower(loc []int, equation string, debug bool) string {
	// get right side of power
	rightSide := equation[strings.Index(equation, "{")+1 : loc[1]-1]
	leftSide := strings.Split(equation[loc[0]:], "^")[0]
	// equation[loc[0]:strings.Index(equation, "^")]
	// if power contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, rightSide)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		rightSide = ParseArgsParen(rightSide, debug)
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
func parseSqrt(loc []int, equation string, debug bool) string {
	// strip sqrt declaration from equation
	innerSqrt := equation[loc[0]+5 : loc[1]-1]
	// if sqrt[] contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, innerSqrt)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		innerSqrt = ParseArgsParen(innerSqrt, debug)
	}
	// if sqrt[] only contains a number, parse it as float
	sqrtFloat, err := strconv.ParseFloat(innerSqrt, 64)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// parse both digits in equation
	equation = equation[:loc[0]] + strconv.FormatFloat(math.Sqrt(sqrtFloat), 'f', -1, 64) + equation[loc[1]:]
	return equation
}
