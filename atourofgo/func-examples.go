
package atourofgo

import "fmt"

/**
*	@note These are exported because they begin with cap
*/

func Add(x int, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

/**
*	@note The third param here is redundant
*/
func Mult(x, y int, z string) int {
	return x*y
}

/**
*	Example of multiple return values; returns both of the given string args
*	in reverse order.
*/
func Swap(x, y string) (string, string) {
	return y, x
}

/**
*	Example of named return types and naked returns.
*
*	Here we return 3 values, each with a name (x, y, and z) and each of type int
*	All we have to do is assign a value to them and call return and they are returned
*	as in the return signature.
*
*	@note The tour says only to naked returns for smaller functions for readabillity.
*/
func Split(sum int) (x, y, z int) {
	x = sum / 3
	y = x / 2
	z = sum - x - y
	return
}


/**
*	Demonstrates using `defer` to defer function calls
*/
func Defer() {

	defer fmt.Println("I'm being defered")
	fmt.Println("I'm not defered")

}
