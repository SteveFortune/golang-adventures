
package main

import (
	"fmt"
	"net/http"
)


type InfoHandler string

func (handler InfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, handler)
}
