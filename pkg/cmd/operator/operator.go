package operator

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Operator is the instance type
type Operator struct {
	Router *mux.Router
}

// NewOperator returns a new Operator instance
func NewOperator() *Operator {
	// return Operator
	return &Operator{
		Router: mux.NewRouter(),
	}
}

// Serve runs the application in server mode
func (o *Operator) Serve() {
	// run listener to daemonize
	o.routes()
	log.Fatal(http.ListenAndServe(":8080", o.Router))
}
