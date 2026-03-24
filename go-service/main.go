package main

import (
	"fmt"
	"log"
	"net/http"

	"go-service/config"
	"go-service/handlers"
	"go-service/pdf"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	
	// Initialize PDF generator
	pdfGenerator := pdf.NewGenerator()
	
	// Create server
	server := handlers.NewServer(pdfGenerator)
	
	// Setup routes
	router := server.SetupRoutes()
	
	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting Go Student Service on port %s", cfg.Port)
	log.Printf("Available endpoints:")
	log.Printf("  GET /health - Health check")
	log.Printf("  POST /students/pdf - Generate PDF report from payload data")
	
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}