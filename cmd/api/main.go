package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"reading-list/internal/data"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

const (
    HOST = "localhost"
    PORT = 5432
)

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")
	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("POSTGRESQL_URL"), "PostgreSQL DSN")
	flag.Parse()

	cfg.dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	HOST, PORT, "postgres", "bucketeer_pass", "bucketeer_db")
	// cfg.dsn = "postgres://postgres:mysecretpassword@localhost/readinglist?sslmode=disable"
	// postgres://postgres:mysecretpassword@localhost/readinglist?sslmode=disable
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
