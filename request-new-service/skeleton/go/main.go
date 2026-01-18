package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from DevOps Team! Your new Go service is ready to be developed.")
	})
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
