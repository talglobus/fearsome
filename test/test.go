// test contains a number of internal tools used only for testing purposes
package test

import "fmt"

// IsStringer() offers an easy way to check if a type implements the `Stringer` interface
func IsStringer(a interface{}) bool {
	_, ok := a.(fmt.Stringer)
	return ok
}
