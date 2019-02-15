// converts input equation to LaTeX format
package calclisrc

import (
	"regexp"
)

func ConvertToLaTeX(equation string) string {
	// find all left and right braces in equation
	braceLocLeft := regexp.MustCompile(`\{`).FindAllStringIndex(equation, -1)
	braceLocRight := regexp.MustCompile(`\}`).FindAllStringIndex(equation, -1)
	// Loop over braces. replace pairs with spaces if they are not sqrt or powers
	for i, braceLoc := range braceLocLeft {
		if string(equation[braceLoc[0]-1]) != "t" && string(equation[braceLoc[0]-1]) != "^" {
			equation = equation[:braceLoc[0]] + " " + equation[braceLoc[1]:]
			equation = equation[:braceLocRight[i][0]] + " " + equation[braceLocRight[i][1]:]
		}
	}
	// find operators that don't have a leading \ character. add the character
	sqSinCosTanRegex := regexp.MustCompile(`([^\\]{1}(sqrt|sin|cos|tan))`)
	for loc := sqSinCosTanRegex.FindStringIndex(equation); loc != nil; loc = sqSinCosTanRegex.FindStringIndex(equation) {
		equation = equation[:loc[0]+1] + "\\" + equation[loc[0]+1:]
	}
	// return the formatted equation
	return equation
}
