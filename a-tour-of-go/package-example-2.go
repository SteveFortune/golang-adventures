
package atourofgo

/**
*	Uses individual import statements...
*/
import "fmt"
import "math/rand"

func PackageExample2() {
	rand.Seed(123)
	fmt.Printf("How you have %g problems", rand.Intn(10))
}
