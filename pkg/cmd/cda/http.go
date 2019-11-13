package cda

import (
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/go-playground/webhooks.v5/github"
)

func (a *Agent) routes() {
	a.Router.HandleFunc("/github", a.githubHandler)
	a.Router.HandleFunc("/healthz", a.healthHandler)
}

func (a *Agent) healthHandler(w http.ResponseWriter, r *http.Request) {
	// respond OK
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}

func (a *Agent) githubHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("github.handler:", r.Method, "request from ", r.RemoteAddr)

	// get the Secret from environment
	secret := os.Getenv("GITHUB_SECRET")

	// Create a new hook config (with secret)
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		handleError(err, w)
		return
	}

	// parse the hooks payload
	payload, err := hook.Parse(r, github.Event("push"))
	if err != nil {
		handleError(err, w)
		return
	}
	pushMessage := payload.(github.PushPayload)

	log.Println("Got Payload from:", pushMessage.Repository)

	// respond OK
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "{}")
}

// error handling function
func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
	log.Printf("error: %s", err.Error())
}
