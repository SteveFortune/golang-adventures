
package main

import (
	"fmt"
)

type Message func (logger BirthdayLogger)

func HappyBirthdayToYou(logger BirthdayLogger) {
	logger.Log("Happy birthday to you\n")
}

func HappyBirthdayDearBen(logger BirthdayLogger) {
	logger.Log(fmt.Sprintf("Happy birthday dear Ben\n"))
}
