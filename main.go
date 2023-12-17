package main

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log information about the request
		fmt.Printf("[%s] %s %s\n", r.Method, r.RequestURI, time.Since(startTime))
	})
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Attach the middleware to the main handler
	mainHandler := http.HandlerFunc(ReqHandler)
	mux.Handle("/", LoggingMiddleware(mainHandler))

	// Start the server with the ServeMux
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
