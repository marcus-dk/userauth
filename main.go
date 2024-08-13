package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"userauth/routes"
)

func main() {
	// Set up the router
	router := routes.SetupRouter()

	// Define the server with the router
	server := &http.Server{
		Addr:    ":8080",  // Address where the server listens
		Handler: router,   // The router that handles incoming requests
	}

	// Run the server in a goroutine so it doesn't block the shutdown handling
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	log.Println("Server running on port 8080")

	// Channel to listen for OS interrupt signals
	quit := make(chan os.Signal, 1)
	// Notify the channel on receiving SIGINT (Ctrl+C) or SIGTERM
	signal.Notify(quit, os.Interrupt)

	// Block until a signal is received
	<-quit
	log.Println("Shutting down server...")

	// Context with a timeout to ensure the server shuts down within a certain period
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
