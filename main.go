package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/callback", callbackHandler)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// MATTR wants all errors to be 404
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read body", http.StatusNotFound)
			return
		}
		defer r.Body.Close()

		if len(body) == 0 {
			http.Error(w, "Content not found", http.StatusNotFound)
			return
		}

		fmt.Println(string(body))

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		//fmt.Fprintf(w, "%s", body)
	} else {
		http.Error(w, "Invalid request method", http.StatusNotFound)
	}
}
