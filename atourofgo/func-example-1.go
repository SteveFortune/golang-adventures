
package atourofgo

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
