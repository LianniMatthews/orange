// Filename cmd/api/main.go
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Global variable to hold application version
const version = "1.0.0"

// Struct for server info
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

// Dependency Injection
type application struct {
	config config
	logger *log.Logger
}

// main Function
func main() {
	var cfg config

	//Get arguments for server configuration
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production")
	flag.StringVar(&cfg.db.dsn, "dsn", os.Getenv("LIME_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	//New logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//open database connection pool
	db, err := openBD(cfg)

	if err != nil {
		logger.Fatal(err)
	}

	//close database connection pool
	defer db.Close()

	//Object of type application
	app := &application{
		config: cfg,
		logger: logger,
	}

	//Server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,      //inactive connections
		ReadTimeout:  10 * time.Second, //time to read request body or header
		WriteTimeout: 10 * time.Second,
	}

	//Start Server
	logger.Printf("starting %s server on port %d", cfg.env, cfg.port)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}

// database connection
func openBD(cfg config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.db.dsn)

	if err != nil {
		return nil, err
	}

	//context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//ping database
	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}
