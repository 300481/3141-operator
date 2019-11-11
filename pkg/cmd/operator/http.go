package operator

import (
	"io"
	"log"
	"net/http"
)

func (o *Operator) routes() {
	o.Router.HandleFunc("/", o.hookHandler)
}

func (o *Operator) hookHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hook.handler:", r.Method, "request from ", r.RemoteAddr)
	io.WriteString(w, "")
}
