//Filename cmd/api/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//Global variable to hold application version
const version = "1.0.0"

//Struct for server info
type config struct {
	port int 
	env string
}

//Dependency Injection
type application struct{
	config config 
	logger *log.Logger
}

//main Function
func main() {
	var cfg config
	
	//Get arguments for server configuration
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production")
	flag.Parse()

	//New logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	
	//Object of type application
	app := &application{
		config: cfg,
		logger: logger,
	}

	//Route
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	//Server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: mux,
		IdleTimeout: time.Minute, //inactive connections
		ReadTimeout: 10 * time.Second, //time to read request body or header
		WriteTimeout: 10 * time.Second,
	}

	//Start Server
	logger.Printf("starting %s server on port %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)
	
}
