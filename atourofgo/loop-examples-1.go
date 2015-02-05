
package atourofgo

import "fmt"

func LoopTo(dst int) {
	for i := 0; i < dst; i++ {
		fmt.Println(i)
	}
}

func LoopFrom(dst int) {
	for dst >= 0 {
		dst--;
		fmt.Println(dst)
	}
}

func LoopForever() {
	for {}
}
