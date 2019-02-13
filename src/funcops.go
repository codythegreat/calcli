// recieves a function operator from parse.go, resolves it to a single float value, returns it
package calclisrc

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func parseSin(loc []int, equation string) string {
	// strip sqrt declaration from equation
	innerSin := equation[loc[0]+4 : loc[1]-1]
	// if sqrt[] contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, innerSin)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		innerSin = parseArgsParen(innerSin)
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
func parseTan(loc []int, equation string) string {
	// strip sqrt declaration from equation
	innerTan := equation[loc[0]+4 : loc[1]-1]
	// if sqrt[] contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, innerTan)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		innerTan = parseArgsParen(innerTan)
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
func parseCos(loc []int, equation string) string {
	// strip sqrt declaration from equation
	innerCos := equation[loc[0]+4 : loc[1]-1]
	// if sqrt[] contains other operators, parse their values:
	match, err := regexp.MatchString(`[^\d\.]*`, innerCos)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if match {
		innerCos = parseArgsParen(innerCos)
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
