
package main

import (
	"fmt"
)

func main() {

	logger := PrintlnBirthdayLogger{}
	queue := &MessageQueue{}

	queue.PushLogger(logger)
	queue.PushMessage(HappyBirthdayToYou)
	queue.PushMessage(HappyBirthdayToYou)
	queue.PushMessage(HappyBirthdayDearBen)
	queue.PushMessage(HappyBirthdayToYou)

	fmt.Printf("Processing birthday queue...\n")
	queue.Process()
}
