
package atourofgo

import "fmt"

/**
*	Declares 3 vars:
*
*	- variable1: an integer variable without a value
*	- variable2: a variable without a type initialiser takes the type
*	  of the value assigned to it, in this case a list of strings
*	- variable3: a variable without a type specifier or value
*/
// For some reason this causes an error...
//var variable1 int, variable2, variable3 = "I am a list", "of strings"


/**
*	Example of short assignment. We can use the ':=' operator to assign
*	a local variable inside of a function with an implicity type.
*/
func PrintLocalVariable(){

	localVar := "I am a local variable, short assignment in action."

	fmt.Println(localVar)

}
