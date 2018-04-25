package handlers

import (
	"log"
	"net/http"

	"github.com/iammallik/sample-heroku-go/app/config"
)

// Liveness handles calls to the '/liveness' endpoint and checks that the
// application is running. Failure of this check will cause Kubernetes to
// restart the application pod.
func Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// Readiness handles calls to the '/readiness' endpoint and checks that the
// application is ready to accept traffic. Failure of this check will cause
// Kubernetes to stop sending traffic to the application.
func Readiness(w http.ResponseWriter, r *http.Request) {
	if !config.CheckEnv() {
		log.Println("Readiness() Err in Env vars due to some being unset")
		http.Error(w, "Environment variables are not set", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
