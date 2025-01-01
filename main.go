package main
import (
	"fmt"
	"net/http"
)

/*
This is the function that will be called when the user visits the page. Handlefunc takes two arguments, the path and the function to be called.
Takes two arguments, the response writer and the request.
*/
func hello(w http.ResponseWriter, r *http.Request) {
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// This is the file server that will serve the files in the current directory.
	// Better practice is to serve files from a specific directory (typically name this "static").
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.HandleFunc("/hello", hello)

	fmt.Println("Starting server on 192.168.1.192:8888")
	http.ListenAndServe("192.168.1.192:8888", nil)
}