package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Serj1c/user-balance/pkg/handler"
	"github.com/Serj1c/user-balance/pkg/users"
	"github.com/Serj1c/user-balance/pkg/util"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Error while reading config file: ", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error while sql.Open: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to db, err: %v\n", err)
	}

	ur := users.NewRepo(db)

	sm := mux.NewRouter()

	uh := handler.NewUserHandler(ur)

	sm.HandleFunc("/deposit", uh.Deposit).Methods("POST")
	sm.HandleFunc("/withdraw", uh.Withdraw).Methods("POST")
	sm.HandleFunc("/transfer", uh.Transfer).Methods("POST")
	sm.HandleFunc("/balance", uh.GetBalance).Methods("GET")
	sm.HandleFunc("/operations", uh.ListAllOperations).Methods("GET")

	server := &http.Server{
		Addr:         config.ServerPort,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		fmt.Printf("Server is listening on port%s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting server: %s", err)
			os.Exit(1)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	fmt.Println("Command to terminate received, shutdown", sig)

	timeoutContext, finish := context.WithTimeout(context.Background(), 30*time.Second)
	defer finish()
	server.Shutdown(timeoutContext)
}
