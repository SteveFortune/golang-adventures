
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

/**
*	An array of strings
*/
var array [12]string

/**
*	Example function of assigning values at array indeces
*/
func ArrayExample() {
	array[0] = "first string value"
	array[1] = "second string value"
}

/**
*	A slice of integers
*/
var slice []int = []int{1, 2, 3}

/**
*	Hey look ! The zero value for a slice is nil.
*/
var aNilSlice []int = nil;

/**
*	Example function that iterates over a slice
*/
func IterateSlice(sliceIter []string) {
	for i := 0; i < len(sliceIter); i++ {
		fmt.Printf(sliceIter[i])
	}
}

/**
*	Exmaple function that slices a slice in half
*/
func SliceInHalf(sliceMe []int) ([]int, []int) {
	length := len(sliceMe)
	halfWay := length/2
	return sliceMe[0: halfWay], sliceMe[halfWay: length]
}

/**
*	Creates an int slice with 4 0-d values and a capacity of 7
*/
func MakeMeASlice() []int {
	return make([]int, 4, 7)
}


/**
*	Exemplifies appendation
*/
func AddToSlice(appendToMe []int) []int {

	appendToMe = append(appendToMe, 123)
	appendToMe = append(appendToMe, 1, 2, 3)

	return appendToMe
}

/**
*	Use range syntax to loop over the given array several times
*/
func RangeLoopMe(loopMe []string) {
	for i, value := range loopMe {
		fmt.Printf("%d, %q", i, value)
	}
	for i := range loopMe {
		fmt.Printf("%d", i)
	}
	for _, value := range loopMe {
		fmt.Printf("%q", value)
	}
}


/**
*	Map example
*/
var dictionary = map[string] struct {
	Property int
	AnotherProperty string
}{
	"First thing": {
		Property: 1,
		AnotherProperty: "thing",
	},
	"Second thing": {
		Property: 2,
		AnotherProperty: "another thing",
	},
}
