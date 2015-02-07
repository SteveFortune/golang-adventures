
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
	if item < rand.Intn(2) {
		return false
	}
	return true
}

/**
*	Redundant, obselete function that demonstrates a simple if statement with
*	a short statement.
*/
func IfExample2(item float64) bool {
	if pow := int(math.Pow(item, 2)); pow < rand.Intn(2) {
		return false
	} else {
		fmt.Printf("Pow %q", pow)
	}
	return true
}


/**
*	Another haleriously redundant example function demonstrating switch syntax
*/
func SwitchExample1(item float64) {

	switch pow := int(math.Pow(item, 2)) ; pow {
		case 4:
			fmt.Println("Its four !")
		case 9:
			fmt.Println("Its nine !")
		default :
			fmt.Println("2 big 2 handle.")
	}

}
