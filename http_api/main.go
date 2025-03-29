package main

import (
	"encoding/json"
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

type postgreetReq struct {
	Name string `json:"name"`
}

func PostGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var reqBody postgreetReq

		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
		}
		if reqBody.Name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello from post body %s", reqBody.Name)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/greet", HelloHandler)
	http.HandleFunc("/api/postgreet", PostGreet)
	fmt.Println("server is running on port 3000")
	// Handle error properly
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
