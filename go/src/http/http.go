package http

import (
    "fmt"
    "net/http"
)
    
type String string

type Hello struct {

}

func (h Hello) ServeHTTP (
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func Main() {
	var h Hello
	http.Handle("/string", h)

	http.ListenAndServe("localhost:8888", nil)
}
