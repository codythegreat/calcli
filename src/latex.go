// converts input equation to LaTeX format
package calclisrc

import (
	"regexp"
)

func ConvertToLaTeX(equation string) string {
	// add extra space to front/end to avoid index out of range
	// define regular expression for input that needs to change
	// for the above regex, simply add a \ character before declaration
	// sin, cos and tan must not have braces
	var bracesToChange []bool
	for _, braceLoc := range regexp.MustCompile(`\{`).FindAllStringIndex(equation, -1) {
		if string(equation[braceLoc[0]-1]) != "t" && string(equation[braceLoc[0]-1]) != "^" {
			bracesToChange = append(bracesToChange, true)
			equation = equation[:braceLoc[0]] + " " + equation[braceLoc[1]:]
		} else {
			bracesToChange = append(bracesToChange, false)
		}
	}
	for i, braceLoc := range regexp.MustCompile(`\}`).FindAllStringIndex(equation, -1) {
		if bracesToChange[i] == true {
			equation = equation[:braceLoc[0]] + " " + equation[braceLoc[1]:]
		}
	}
	sqSinCosTanRegex := regexp.MustCompile(`(sqrt|sin|cos|tan)`)
	for loc := sqSinCosTanRegex.FindStringIndex(equation); loc != nil; loc = sqSinCosTanRegex.FindStringIndex(equation) {
		equation = equation[:loc[0]] + "\\" + equation[loc[0]:]
	}
	// return the formatted equation
	return equation
}
