package main

import (
    "log"
    "net/http"
    "os"
    "catgpt-backend/pkg/handler"
    "catgpt-backend/pkg/middleware"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    router := mux.NewRouter()
    router.Use(middleware.LoggingMiddleware)

    router.HandleFunc("/chat", handler.ChatHandler).Methods("POST")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}