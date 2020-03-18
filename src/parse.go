// uses regex to parse operations in the equation
package calclisrc

import (
	"regexp"
)

// parse all PEMDAS operations
func parseSimpleOperations(equation string, debug bool) string {
	// simple operations:
	addOpRegex := regexp.MustCompile(`\-?\d+(\.\d*)?\+\-?\d+(\.\d*)?`)
	subOpRegex := regexp.MustCompile(`\-?\d+(\.\d*)?\-\-?\d+(\.\d*)?`)
	multOpRegex := regexp.MustCompile(`\-?\d+(\.\d*)?\*\-?\d+(\.\d*)?`)
	divOpRegex := regexp.MustCompile(`\-?\d+(\.\d*)?/\-?\d+(\.\d*)?`)

	// parse all multiplication in equation
	for multOpIndex := multOpRegex.FindStringIndex(equation); multOpIndex != nil; multOpIndex = multOpRegex.FindStringIndex(equation) {
		equation = parseMult(multOpIndex, equation, debug)
	} // parse all division in equation
	for divOpIndex := divOpRegex.FindStringIndex(equation); divOpIndex != nil; divOpIndex = divOpRegex.FindStringIndex(equation) {
		equation = parseDiv(divOpIndex, equation, debug)
	}
	// parse all addition in equation
	for addOpIndex := addOpRegex.FindStringIndex(equation); addOpIndex != nil; addOpIndex = addOpRegex.FindStringIndex(equation) {
		equation = parseAdd(addOpIndex, equation, debug)
	}
	// parse all subtraction in equation
	for subOpIndex := subOpRegex.FindStringIndex(equation); subOpIndex != nil; subOpIndex = subOpRegex.FindStringIndex(equation) {
		equation = parseSub(subOpIndex, equation, debug)
	}
	return equation
}

// parses all advanced operations (any that have {} brackets)
func parseAdvancedOperations(equation string, debug bool) string {
	// starts from the innermost power in the equation
	powerOpRegex := regexp.MustCompile(`\-?\d+(\.\d+)?\^\{[^\{\}]*\}`)
	// starts from the innermost sqrt in the equation
	sqrtOpRegex := regexp.MustCompile(`sqrt\{[^\{\}]+\}`)
	// starts from the innermost sin in the equation
	sinOpRegex := regexp.MustCompile(`sin\{[^\{\}]+\}`)
	// starts from the innermost cos in the equation
	cosOpRegex := regexp.MustCompile(`cos\{[^\{\}]+\}`)
	// starts from the innermost tan in the equation
	tanOpRegex := regexp.MustCompile(`tan\{[^\{\}]+\}`)

	// parse all sqrt in equation
	for sqrtOpIndex := sqrtOpRegex.FindStringIndex(equation); sqrtOpIndex != nil; sqrtOpIndex = sqrtOpRegex.FindStringIndex(equation) {
		equation = parseSqrt(sqrtOpIndex, equation, debug)
	}
	// parse all exponents in equation
	for powerOpIndex := powerOpRegex.FindStringIndex(equation); powerOpIndex != nil; powerOpIndex = powerOpRegex.FindStringIndex(equation) {
		equation = parsePower(powerOpIndex, equation, debug)
	}
	// parse all sin in equation
	for sinOpIndex := sinOpRegex.FindStringIndex(equation); sinOpIndex != nil; sinOpIndex = sinOpRegex.FindStringIndex(equation) {
		equation = parseSin(sinOpIndex, equation, debug)
	}
	// parse all cos in equation
	for cosOpIndex := cosOpRegex.FindStringIndex(equation); cosOpIndex != nil; cosOpIndex = cosOpRegex.FindStringIndex(equation) {
		equation = parseCos(cosOpIndex, equation, debug)
	}
	// parse all tan in equation
	for tanOpIndex := tanOpRegex.FindStringIndex(equation); tanOpIndex != nil; tanOpIndex = tanOpRegex.FindStringIndex(equation) {
		equation = parseTan(tanOpIndex, equation, debug)
	}
	return equation
}

// return equation with all simple and advanced operations parsed
func parseArgs(equation string, debug bool) string {
	return parseSimpleOperations(parseAdvancedOperations(equation, debug), debug)
}
