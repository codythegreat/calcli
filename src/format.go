// performs formatting fixes and adjustments on equation
package calclisrc

import (
	"regexp"
	"strings"
)

func RemoveSpacesFromEquation(args []string) string {
	if len(args) > 2 {
		return regexp.MustCompile(` `).ReplaceAllString(strings.Join(args[2:len(args)], ""), "")
	}
	return regexp.MustCompile(` `).ReplaceAllString(strings.Join(args[1:len(args)], ""), "")
}
