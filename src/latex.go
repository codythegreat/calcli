// converts input equation to LaTeX format
package calclisrc

import (
	"fmt"
	"regexp"
	"strings"
)

func ConvertToLaTeX(equation string) string {
	// define regular expression for input that needs to change
	sqrtRegex, err := regexp.Compile(`sqrt`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	sinCosTanRegex, err := regexp.Compile(`(sin|cos|tan)\{[^\{\}]*\}`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// for sqrt, simply add a \ character before declaration
	equation = sqrtRegex.ReplaceAllString(equation, "\\sqrt")
	// for sin, cos, and tan: start from inner function, add \ character and replace { } with spaces
	for b := sinCosTanRegex.MatchString(equation); b != false; b = sinCosTanRegex.MatchString(equation) {
		for sinCosTanIndex := sinCosTanRegex.FindStringIndex(equation); sinCosTanIndex != nil; sinCosTanIndex = sinCosTanRegex.FindStringIndex(equation) {
			convertedString := "\\" + strings.Replace(strings.Replace(equation[sinCosTanIndex[0]:sinCosTanIndex[1]], "{", " ", -1), "}", " ", -1)
			equation = equation[:sinCosTanIndex[0]] + convertedString + equation[sinCosTanIndex[1]:]
		}
	}
	// return the formatted equation
	return equation
}
