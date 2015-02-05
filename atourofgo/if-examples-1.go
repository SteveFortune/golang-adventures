
package atourofgo

import (
	"fmt"
	"math"
	"math/rand"
)

/**
*	Redundant, obselete function that demonstrates a simple if statement
*/

func IfExample1(item int) bool {
	if item < rand.Intn() {
		return false
	}
	return true
}

/**
*	Redundant, obselete function that demonstrates a simple if statement with
*	a short statement.
*/
func IfExample2(item int) bool {
	if pow := math.Pow(2); pow < rand.Intn() {
		return false
	}
	else{
		fmt.Printf('Pow %q', pow)
	}
	return true
}
