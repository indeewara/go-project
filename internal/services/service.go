package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-project/internal/handlers"
)

func RunService() {
	http.HandleFunc("/convert", handlers.ConvertHandler)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
