
package atourofgo

import (
	"math/cmplx"
)

/**
*	Block variable declaration; variable assignments of different types
*/
var (
	var1 bool = false
	var2 uint32 = 1 << 32 - 1
	var3 complex128 = cmplx.Sqrt(-5 + 12i)
	var4 complex64 = 21i
)

/**
*	Initialised to its 'zero' value, an empty string
*/
var var5 string

/**
*	Example of type convertion, converts an int into a float64 and string
*	respectively
*/
func Convert(integer int) (float64, string) {
	floatValue := float64(integer)
	stringValue := string(integer)
	return floatValue, stringValue
}
