package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Declare an instance of the config struct.
	var cfg config

	// syntax  flag.IntVar(p *int, name string, value int, usage string)
	// your storing the cfg.env value in pointer p
	flag.IntVar(&cfg.port, "port", 4000, "API server port")

	// syntax  flag.StringVar(p *string, name string, value string, usage string)
	// your storing the cfg.env value in pointer p
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// We use &application{} because we want app to be a pointer to an application struct.
	// That’s so we can later pass app around efficiently (not copy the whole struct) and modify its contents inside methods.
	// and logger doesn't need an & because log.New returns a pointer (check through intellicence)
	// TLDR
	//Use & when the function needs a pointer and you only have a value.
	//Don’t use & if you already have a pointer returned by something like log.New.
	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server running on port %d\nVersion: %s\n", app.config.port, version)
	})
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	// Declare a HTTP server with some sensible timeout settings, which listens on the
	// port provided in the config struct and uses the servemux we created above as the
	// handler.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()

	logger.Fatal(err)

}
