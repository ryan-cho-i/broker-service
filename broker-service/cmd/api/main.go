// go mod init broker
package main

import (
	"fmt"
	"log"
	"net/http"
)

// go get github.com/go-chi/chi/v5
// go get github.com/go-chi/chi/v5/middleware
// go get github.com/go-chi/cors

const webPort = "80"

type Config struct {}

func main() {
	
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic((err))
	}
}	