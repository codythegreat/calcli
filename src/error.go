// input verification functions. If these return true main will handle the error
package calclisrc

import (
	"regexp"
	"strings"
)

// use a regex of each acceptable symbol/character
// find longest match and verify it matches the input
func VerifyEquationHasProperSyntax(equation string) bool {
	syntax := regexp.MustCompile(`(\.|\+|\-|\*|\/|\d|\{|\}|\^|sin|cos|tan|sqrt|\(|\))+`)
	syntax.Longest()
	return syntax.FindString(equation) == equation
}

func VerifyInputHasEquation(args int) bool {
	return args > 1
}

func VerifyEquationHasProperBrackets(equation string) bool {
	return regexp.MustCompile(`(\^|sqrt|sin|cos|tan)[^\{]`).FindIndex([]byte(equation)) != nil || strings.Count(equation, "{") != strings.Count(equation, "}")
}
