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

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
		}

		response := map[string]string{
			"message": fmt.Sprintf("hi this is a  response %s i know ur mail is %s", user.Name, user.Email),
		}

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "invalid req", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/greet", HelloHandler)
	http.HandleFunc("/api/postgreet", PostGreet)
	http.HandleFunc("/user", jsonHandler)
	fmt.Println("server is running on port 3000")
	// Handle error properly
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
