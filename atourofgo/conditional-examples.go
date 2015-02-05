
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


/**
*	Another haleriously redundant example function demonstrating switch syntax
*/
func SwitchExample1(item int) string{

	switch pow = math.Pow(2) ; pow {
		case 4:
			fmt.Println("Its four !")
		case 9:
			fmt.Println("Its nine !")
		default :
			fmt.Println("2 big 2 handle.")
	}

}
