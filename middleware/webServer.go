package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

// ErrMsgHandler defines the error message handler
func ErrMsgHandler(msg string, err error) {
	if err != nil {
		log.Println(msg, err.Error())
		os.Exit(1)
	}
}

// InitializeWebServer starts the webserver by
// - taking in a `port` as string
// - calling the net/http package to initialize a webserver listening at the define `port`
// - using package errors to catch and handle service startup issues
func InitializeWebServer(port string) {
	// route handler
	routeHandler := NewRouter()
	http.Handle("/", routeHandler)

	// http server
	log.Println("==== ==== ==== ====")
	log.Println("Starting webserver on port: " + port)
	log.Fatal(errors.Wrap(http.ListenAndServe(":"+port, nil), "Failed to start webserver at port:"+port))
}
