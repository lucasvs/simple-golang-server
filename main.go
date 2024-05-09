package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/user/{name}", UserHandler).Methods("GET")

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      r,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
    }

    go func() {
        log.Println("Starting the server on :8080")
        if err := srv.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down the server...")

    // gracefully shutdown with delay here
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server shutdown error: %v", err)
    }
    log.Println("Server gracefully stopped")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi change this count: 1")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    fmt.Fprintf(w, "Hello, %s!", name)
}