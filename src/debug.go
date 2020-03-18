// prints function calls, inputs and outputs
package calclisrc

import (
	"fmt"
	"reflect"
	"runtime"
)

// helper function to return argument function's name as string
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// debug message at start of function
// function name starts at 35th index
func DebugMessageFunctionStart(function interface{}, output string) {
	fmt.Printf("DEBUGGER: %-20v called with %s as input\n", getFunctionName(function)[35:], output)
}

// debug message at end of function
// function name starts at 35th index
func DebugMessageFunctionEnd(function interface{}, output string) {
	fmt.Printf("DEBUGGER: %-20v returned with %s as ouput\n", getFunctionName(function)[35:], output)
}
