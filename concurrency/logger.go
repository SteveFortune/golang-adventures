
package main

import "fmt"

type BirthdayLogger interface {
	Log(msg string)
}

type PrintlnBirthdayLogger struct {}

func (logger PrintlnBirthdayLogger) Log(msg string){
	fmt.Printf(msg)
}
