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
	sinRegex, err := regexp.Compile(`sin\{.*\}`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	cosRegex, err := regexp.Compile(`cos\{.*\}`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	tanRegex, err := regexp.Compile(`tan\{.*\}`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// for sqrt, simply add a \ character before declaration
	equation = sqrtRegex.ReplaceAllString(equation, "\\sqrt")
	// all other functions will be stripped of "{" and "}" with a \ prepended
	for _, index := range sinRegex.FindAllStringIndex(equation, -1) {
		convertedString := "\\" + strings.Replace(strings.Replace(equation[index[0]:index[1]], "{", " ", -1), "}", " ", -1)
		equation = equation[:index[0]] + convertedString + equation[index[1]:]
	}
	for _, index := range cosRegex.FindAllStringIndex(equation, -1) {
		convertedString := "\\" + strings.Replace(strings.Replace(equation[index[0]:index[1]], "{", " ", -1), "}", " ", -1)
		equation = equation[:index[0]] + convertedString + equation[index[1]:]
	}
	for _, index := range tanRegex.FindAllStringIndex(equation, -1) {
		convertedString := "\\" + strings.Replace(strings.Replace(equation[index[0]:index[1]], "{", " ", -1), "}", " ", -1)
		equation = equation[:index[0]] + convertedString + equation[index[1]:]
	}
	// return the formatted equation
	return equation
}
