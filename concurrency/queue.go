
package main

import "time"

type MessageQueue struct {
	messages []Message
	loggers []BirthdayLogger
}


func (queue *MessageQueue) PushMessage(msg Message) {
	queue.messages = append(queue.messages, msg)
}


func (queue *MessageQueue) PushLogger(logger BirthdayLogger) {
	queue.loggers = append(queue.loggers, logger)
}


func (queue *MessageQueue) Process() {
	go func () {
		for _, msg := range queue.messages {
			for _, logger := range queue.loggers {
				msg(logger)
				time.Sleep(1000*time.Millisecond)
			}
		}
	}()
}
