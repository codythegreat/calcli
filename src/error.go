// checks input equation for correct formatting, and
// returns bool
package calclisrc

import (
	"regexp"
	"strings"
)

func VerifyInputHasEquation(args int) bool {
	return args > 1
}

func VerifyEquationHasProperBrackets(equation string) bool {
	return regexp.MustCompile(`(\^|sqrt|sin|cos|tan)[^\{]`).FindIndex([]byte(equation)) != nil || strings.Count(equation, "{") != strings.Count(equation, "}")
}
