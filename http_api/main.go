package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		name := r.URL.Query().Get("name")
		fmt.Fprintf(w, "hello %s", name)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	http.HandleFunc("/greet", HelloHandler)
	fmt.Println("server is running on port 3000")
	http.ListenAndServe(":3000", nil)
	// Handle error properly
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
