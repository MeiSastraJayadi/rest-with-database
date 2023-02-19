package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MeiSastraJayadi/rest-with-datatabase/deliver"
	"github.com/MeiSastraJayadi/rest-with-datatabase/repository"
	_ "github.com/go-sql-driver/mysql"
	gorillahandler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
)

var Address = env.String("ADDRESS", false, "localhost:9090", "Port Address")
var connection = env.String("CONNECTION", false, "root:@tcp(localhost:3306)/golang-db", "Connection to database")
var driver = env.String("DRIVER", false, "mysql", "Use mysql driver")
var maxConnection = env.Int("MAX_CONNECTION", false, 100, "Max connection to database")
var idleConnection = env.Int("IDLE_CONNECTION", false, 20, "Minimum idle connection")
var idleTime = env.Duration("IDLE_DURATION", false, time.Minute, "Idle time")
var lifeTime = env.Duration("LIFETIME", false, 30*time.Minute, "Life time of a connection")

func main() {
	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "e-commerce",
			Level: hclog.LevelFromString("DEBUG"),
		},
	)
	envErr := env.Parse()
	if envErr != nil {
		l.Error("Error when read env", "err", envErr)
		os.Exit(1)
	}
	db_conn := repository.NewConnection(l, *connection, *driver, *maxConnection, *idleConnection, *idleTime, *lifeTime)
	db, err := db_conn.CreateConnection()
	defer db.Close()
	if err != nil {
		l.Error("Error when connect to database", "error", err)
		os.Exit(1)
	}
	categoryDeliver := deliver.NewCategoryDeliver(db, l)
	ownerDeliver := deliver.NewOwnerDeliver(db, l)

	rt := mux.NewRouter()
	//Category Handler
	cg := rt.Methods(http.MethodGet).Subrouter()
	cg.HandleFunc("/category", categoryDeliver.GetAll)

	cpost := rt.Methods(http.MethodPost).Subrouter()
	cpost.HandleFunc("/category", categoryDeliver.Create)

	cdel := rt.Methods(http.MethodDelete).Subrouter()
	cdel.HandleFunc("/category/{id:[0-9]+}", categoryDeliver.Delete)

	cput := rt.Methods(http.MethodPut).Subrouter()
	cput.HandleFunc("/category/{id:[0-9]+}", categoryDeliver.Update)

	//Owner Handler
	og := rt.Methods(http.MethodGet).Subrouter()
	og.HandleFunc("/owner", ownerDeliver.Fetch)

	opost := rt.Methods(http.MethodPost).Subrouter()
	opost.HandleFunc("/owner", ownerDeliver.Create)

	delown := rt.Methods(http.MethodDelete).Subrouter()
	delown.HandleFunc("/owner/{id:[0-9]+}", ownerDeliver.Delete)

	putown := rt.Methods(http.MethodPut).Subrouter()
	putown.HandleFunc("/owner/{id:[0-9]+}", ownerDeliver.Update)

	//Product Handler
	productDeliver := deliver.NewProductDeliver(db, l)

	ppost := rt.Methods(http.MethodPost).Subrouter()
	ppost.HandleFunc("/product", productDeliver.Create)

	pget := rt.Methods(http.MethodGet).Subrouter()
	pget.HandleFunc("/product", productDeliver.FetchAll)

	pdel := rt.Methods(http.MethodDelete).Subrouter()
	pdel.HandleFunc("/product/{id:[0-9]+}", productDeliver.Delete)

	pput := rt.Methods(http.MethodPut, http.MethodPatch).Subrouter()
	pput.HandleFunc("/product/{id:[0-9]+}", productDeliver.Update)

	//Consument Handler
	consumentHandler := deliver.NewConsumentDeliver(db, l)
	conpost := rt.Methods(http.MethodPost).Subrouter()
	conpost.HandleFunc("/consument", consumentHandler.Create)

	conget := rt.Methods(http.MethodGet).Subrouter()
	conget.HandleFunc("/consument", consumentHandler.Fetch)

	condel := rt.Methods(http.MethodDelete).Subrouter()
	condel.HandleFunc("/consument/{id:[0-9]+}", consumentHandler.Delete)

	conput := rt.Methods(http.MethodPut).Subrouter()
	conput.HandleFunc("/consument/{id:[0-9]+}", consumentHandler.Update)

	cors := gorillahandler.CORS(gorillahandler.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         *Address,
		Handler:      cors(rt),
		IdleTimeout:  4 * time.Minute,
		WriteTimeout: time.Minute,
		ReadTimeout:  40 * time.Second,
	}

	go func() {
		l.Info("Starting Starting...")
		err = server.ListenAndServe()
		if err != nil {
			l.Error("Unable to starting server", "error", err)
			os.Exit(1)
		}
	}()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)
	is := <-ch
	if is != nil {
		l.Info("Recieved interrupt signal", "info", is)
	}

	tc, _ := context.WithTimeout(context.Background(), 30*time.Minute)
	server.Shutdown(tc)
}
