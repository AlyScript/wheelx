package main

/*
Author: Adam Aly
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	startTime 		= time.Now()
	requestCount 	int
	logFile 		*os.File
	mu 				sync.Mutex
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

func infoHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime)
	requestCount++
	fmt.Fprintf(w, "Welcome to wheelx!\n Server Start Time: %s\nUptime: %s\nRequest Count: %d\n", startTime.Format(time.RFC3339), uptime, requestCount)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC1123)
    fmt.Fprintf(w, "Current server time: %s\n", currentTime)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    fmt.Fprintf(w, "Total requests received: %d\n", requestCount)
}

func loggingFileServerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received from %s at %s\n", r.RemoteAddr, time.Now().Format((time.RFC3339)))
		next.ServeHTTP(w, r)
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"status": "ok"}
    jsonResponse, _ := json.Marshal(response)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func main() {

	// Define a flag for the port number
	port := flag.String("p", "8888", "Port to run the server on.")
	ip := flag.String("i", "192.168.1.192", "IP address to run the server on.")
	flag.Usage = func() {
		fmt.Printf("\n%s\n%s\n", "Usage: go run main.go -p <port>", "Usage: go run main.go -p <port> -i <ip>")
		flag.PrintDefaults()
	}
	flag.Parse()

	logFile, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	var socket string = *ip + ":" + *port

	// This is the file server that will serve the files in the current directory.
	// Better practice is to serve files from a specific directory (typically name this "static").
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", loggingFileServerHandler(fs))

	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Welcome to wheelx! Starting server on " + socket)
	log.Printf("Starting server on %s\n", socket)
	http.ListenAndServe(socket, nil)
}