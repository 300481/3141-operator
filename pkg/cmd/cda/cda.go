package cda

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Agent is the instance type
type Agent struct {
	Router *mux.Router
}

// NewAgent returns a new Agent instance
func NewAgent() *Agent {
	// return Agent
	return &Agent{
		Router: mux.NewRouter(),
	}
}

// Serve runs the application in server mode
func (a *Agent) Serve() {
	// run listener to daemonize
	a.routes()
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
