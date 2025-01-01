package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

/*
This is the function that will be called when the user visits the page. Handlefunc takes two arguments, the path and the function to be called.
Takes two arguments, the response writer and the request.
*/
// func hello(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Request received from %s at %s\n", r.RemoteAddr, time.Now().Format((time.RFC3339)))

// 	for name, values := range r.Header {
// 		for _, value := range values {
// 			fmt.Printf("%s: %s\n", name, value)
// 		}
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Hello, World!")
// }

func loggingFileServerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received from %s at %s\n", r.RemoteAddr, time.Now().Format((time.RFC3339)))
		next.ServeHTTP(w, r)
	})
}

func main() {

	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// This is the file server that will serve the files in the current directory.
	// Better practice is to serve files from a specific directory (typically name this "static").
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", loggingFileServerHandler(fs))

	// http.HandleFunc("/hello", hello)

	fmt.Println("Starting server on 192.168.1.192:8888")
	http.ListenAndServe("192.168.1.192:8888", nil)
}