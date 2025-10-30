package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"govite/handlers"
)

func main() {
	// Parse command line flags
	isDev := flag.Bool("dev", false, "Run in development mode with Vite dev server")
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	// Determine if we're in development mode (also check environment variable)
	devMode := *isDev || os.Getenv("DEV_MODE") == "true"

	// Set up routes
	http.HandleFunc("/", handlers.HomeHandler(devMode))

	// Serve static files in production mode
	if !devMode {
		fs := http.FileServer(http.Dir("dist"))
		http.Handle("/dist/", http.StripPrefix("/dist/", fs))
	}

	// Serve assets directory
	assetsFS := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assetsFS))

	// Start server
	addr := ":" + *port
	if devMode {
		log.Printf("Starting server in DEVELOPMENT mode on http://localhost%s", addr)
		log.Println("Make sure Vite dev server is running on http://localhost:5173")
		log.Println("Run: npm run dev")
	} else {
		log.Printf("Starting server in PRODUCTION mode on http://localhost%s", addr)
		log.Println("Make sure to build assets first: npm run build")
	}

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(fmt.Errorf("server failed to start: %w", err))
	}
}
