
package main

import (
	"fmt"
	"net/http"
)


type WelcomeHandler struct{
	title string
	greeting string
}

func (handler WelcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, handler.title)
	fmt.Fprint(w, "\n")
	fmt.Fprint(w, handler.greeting)
}
