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
func DebugMessageFunctionStart(function interface{}, output string) {
	fmt.Printf("DEBUGGER: %s called with %s as input\n", getFunctionName(function), output)
}

// debug message at end of function
func DebugMessageFunctionEnd(function interface{}, output string) {
	fmt.Printf("DEBUGGER: %s returned with %s as ouput\n", getFunctionName(function), output)
}
