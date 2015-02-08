
package main

import (
	"fmt"
	"net/http"
)


type AboutHandler string

func (handler AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, handler)
}
