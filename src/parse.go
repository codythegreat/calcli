// uses regex to parse operations in the equation
package calclisrc

import (
	"regexp"
)

func parseArgs(args string) string {
	// variable to be returned
	returnString := args

	// regular expressions to interpret user input:
	addOpRegex := regexp.MustCompile(`\-?\d+(\.\d*)?\+\-?\d+(\.\d*)?`)
	subOpRegex := regexp.MustCompile(`\-?\d+(\.\d*)?\-\-?\d+(\.\d*)?`)
	multOpRegex := regexp.MustCompile(`\-?\d+(\.\d*)?\*\-?\d+(\.\d*)?`)
	divOpRegex := regexp.MustCompile(`\-?\d+(\.\d*)?/\-?\d+(\.\d*)?`)
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

	// while there are functions in the equation (any operator that wraps its rigth value in "{}"):
	for b := strings.Contains(returnString, "{"); b != false; b = strings.Contains(returnString, "{") {
		// parse all sqrt in equation
		for sqrtOpIndex := sqrtOpRegex.FindStringIndex(returnString); sqrtOpIndex != nil; sqrtOpIndex = sqrtOpRegex.FindStringIndex(returnString) {
			returnString = parseSqrt(sqrtOpIndex, returnString)
		}
		// parse all exponents in equation
		for powerOpIndex := powerOpRegex.FindStringIndex(returnString); powerOpIndex != nil; powerOpIndex = powerOpRegex.FindStringIndex(returnString) {
			returnString = parsePower(powerOpIndex, returnString)
		}
		// parse all sin in equation
		for sinOpIndex := sinOpRegex.FindStringIndex(returnString); sinOpIndex != nil; sinOpIndex = sinOpRegex.FindStringIndex(returnString) {
			returnString = parseSin(sinOpIndex, returnString)
		}
		// parse all cos in equation
		for cosOpIndex := cosOpRegex.FindStringIndex(returnString); cosOpIndex != nil; cosOpIndex = cosOpRegex.FindStringIndex(returnString) {
			returnString = parseCos(cosOpIndex, returnString)
		}
		// parse all tan in equation
		for tanOpIndex := tanOpRegex.FindStringIndex(returnString); tanOpIndex != nil; tanOpIndex = tanOpRegex.FindStringIndex(returnString) {
			returnString = parseTan(tanOpIndex, returnString)
		}
	}

	// parse all multiplication in equation
	for multOpIndex := multOpRegex.FindStringIndex(returnString); multOpIndex != nil; multOpIndex = multOpRegex.FindStringIndex(returnString) {
		returnString = parseMult(multOpIndex, returnString)
	} // parse all division in equation
	for divOpIndex := divOpRegex.FindStringIndex(returnString); divOpIndex != nil; divOpIndex = divOpRegex.FindStringIndex(returnString) {
		returnString = parseDiv(divOpIndex, returnString)
	}
	// parse all addition in equation
	for addOpIndex := addOpRegex.FindStringIndex(returnString); addOpIndex != nil; addOpIndex = addOpRegex.FindStringIndex(returnString) {
		returnString = parseAdd(addOpIndex, returnString)
	}
	// parse all subtraction in equation
	for subOpIndex := subOpRegex.FindStringIndex(returnString); subOpIndex != nil; subOpIndex = subOpRegex.FindStringIndex(returnString) {
		returnString = parseSub(subOpIndex, returnString)
	}

	// return string after parsing through order of operations
	return returnString
}
