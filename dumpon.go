package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Declare version as a global variable
var version = "development"

// Declare maxMemory as a global variable
var maxMemory int64

func requestHandler(w http.ResponseWriter, r *http.Request) {
	requestTime := time.Now().Format(time.DateTime)

	fmt.Println(requestTime, "-------------Start-------------")
	fmt.Println(requestTime, "Request URL:", r.Host+r.URL.String())
	fmt.Println(requestTime, "Request Method:", r.Method)
	fmt.Println()

	// Print headers
	fmt.Println(requestTime, "Request Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Println(requestTime, name+": "+value)
		}
	}
	fmt.Println()

	// Print URL parameters
	fmt.Println(requestTime, "URL Parameters:")
	queryParams := r.URL.Query()
	for param, values := range queryParams {
		for _, value := range values {
			fmt.Println(requestTime, param+": "+value)
		}
	}
	fmt.Println()

	// Check if multipart form data
	if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		// Parse the form with the specified memory limit
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			log.Println(requestTime, "Error parsing multipart form:", err)
			http.Error(w, "Unable to parse multipart form", http.StatusBadRequest)
			return
		}

		// Print other form fields, skipping file fields
		fmt.Println(requestTime, "Form Fields:")
		for key, values := range r.MultipartForm.Value {
			if _, found := r.MultipartForm.File[key]; found {
				// Skip fields that are files
				continue
			}
			for _, value := range values {
				fmt.Println(requestTime, key+": "+value)
			}
		}
		fmt.Println()

		// Check for files
		fmt.Println(requestTime, "Files:")
		for key, files := range r.MultipartForm.File {
			for _, file := range files {
				fmt.Println(requestTime, key+": "+file.Filename)
			}
		}
	} else {
		// For non-multipart data, just read the body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(requestTime, "Error reading request body:", err)
		} else {
			fmt.Println(requestTime, "Request Body:")
			fmt.Println(string(body))
		}
	}

	fmt.Println(requestTime, "--------------End--------------")
}

func main() {
	// Parse command-line arguments
	memoryLimit := flag.Int64("m", 10, "Memory limit for parsing multipart form data in megabytes")
	port := flag.Int64("p", 80, "Port to run the server on")

	// Help message
	flag.Usage = func() {
		fmt.Println("Version:", version)
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Convert megabytes to bytes
	maxMemory = int64(*memoryLimit) * 1024 * 1024
	address := fmt.Sprintf(":%d", *port)

	http.HandleFunc("/", requestHandler)
	server := &http.Server{Addr: address}

	// Start server in a goroutine
	go func() {
		fmt.Println("Dumpon v"+version+" started on port", *port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port %s: %v\n", address, err)
		}
	}()

	// Set up signal handler to catch interrupts
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for an interrupt signal
	<-stop
	fmt.Println("\nShutting down the server...")

	// Create a timeout for force server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown the server: %v", err)
	}

	fmt.Println("Server exiting")
}
