
package atourofgo

import (
	"math/cmplx"
	"fmt"
)


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


/**
* Const declarations
*/
const Thing = "I am a constant"
const (
	Thing2 = "I am another const"
	Thing3 = "I am a const declared in a block"
)


/**
*	Pointers and suchlike...
*/
func PointerExample() {

	value := "a string"
	pointer := &value
	*pointer = "changed the string"
}


/**
*	A simple little struct
*/
type AStruct struct {
	IntegerVar int
	StringVar string
}

/**
*	Exemplifies different struct allocation sytanxes and struct pointer
* 	transparency
*/
func StructExample(){

	struct1 := AStruct{1, "string me"}
	struct2 := AStruct{StringVar: "stringme", IntegerVar: 1}
	pointerToStruct1 := &struct1

	fmt.Printf(string(struct1.IntegerVar))
	fmt.Printf(struct2.StringVar)
	fmt.Printf(pointerToStruct1.StringVar)

}