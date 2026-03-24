package main

import (
	"fmt"
	"log"
	"net/http"

	"go-service/config"
	"go-service/handlers"
	"go-service/pdf"
	"go-service/student"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	
	// Initialize services
	studentService := student.NewService(cfg.StudentAPIURL)
	pdfGenerator := pdf.NewGenerator()
	
	// Create server (supporting both external API and payload-based PDF generation)
	server := handlers.NewServer(studentService, pdfGenerator)
	
	// Setup routes
	router := server.SetupRoutes()
	
	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting Go Student Service on port %s", cfg.Port)
	log.Printf("Student API URL: %s", cfg.StudentAPIURL)
	log.Printf("Available endpoints:")
	log.Printf("  GET /health - Health check")
	log.Printf("  GET /students/{id} - Get student details as JSON")
	log.Printf("  GET /students/{id}/pdf - Generate PDF report for student (external API)")
	log.Printf("  POST /students/pdf - Generate PDF report from payload data")
	
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}