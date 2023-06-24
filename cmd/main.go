package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/project/handlers" // Update with your package path
)

func main() {
	http.HandleFunc("/convert", handlers.ConvertHandler)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
