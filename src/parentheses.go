// finds innermost parentheses and sends its contents to parse.go/parseArgs

package calclisrc

import (
	"fmt"
	"regexp"
)

func ParseArgsParen(args string, debug bool) string {
	if debug {
		DebugMessageFunctionStart(ParseArgsParen, args)
	}
	// define parentheses regex (only finds inner parentheses)
	parenOpRegex, err := regexp.Compile(`\([^\(\)]+\)`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// parse all parentheses, starting from innermost and working up to outermost
	for parenIndexes := parenOpRegex.FindStringIndex(args); parenIndexes != nil; parenIndexes = parenOpRegex.FindStringIndex(args) {
		args = args[:parenIndexes[0]] + parseArgs(args[parenIndexes[0]+1:parenIndexes[1]-1], debug) + args[parenIndexes[1]:]
	}
	if debug {
		DebugMessageFunctionEnd(ParseArgsParen, args)
	}
	// Once we know that all parentheses are resolved, parse other args and return string answer
	return parseArgs(args, debug)
}
