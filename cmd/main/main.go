package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/fuckthinkpad/internal/routes"
)

func main() {
	log.Info("Server Starting at :8080")
	http.ListenAndServe(":8080", routes.Routes())
}
