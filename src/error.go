// checks input equation for correct formatting, and
// returns bool
package calclisrc

import (
	"regexp"
	"strings"
)

func DetectInputError(equation string) bool {
	missingCurlyBracket := regexp.MustCompile(`(\^|sqrt|sin|cos|tan)[^\{]`).FindIndex([]byte(equation))
	if missingCurlyBracket != nil || strings.Count(equation, "{") != strings.Count(equation, "}") {
		return true
	}
	return false
}
