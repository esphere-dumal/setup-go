package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// The code is not important, it's just a simple Go server using Gorilla Mux and Logrus for logging.
// We hope to use these code to test setup-go action.

var logger = logrus.New()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Home endpoint hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Home Page!"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("About endpoint hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is a simple Go server using Gorilla Mux and Logrus!"))
}

func main() {
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")

	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatalf("Could not start server: %s", err)
	}
}
