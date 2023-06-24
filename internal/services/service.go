package services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-project/internal/handlers"
	"github.com/patrickmn/go-cache"
)

func RunService() {
	c := cache.New(3*time.Hour, 6*time.Hour)

	myApp := handlers.App{
		Cache: c,
	}

	http.HandleFunc("/convert", myApp.ConvertHandler)

	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
