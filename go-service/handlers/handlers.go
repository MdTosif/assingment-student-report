package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go-service/pdf"
	"go-service/student"

	"github.com/gorilla/mux"
)

// Server represents the HTTP server
type Server struct {
	pdfGenerator *pdf.Generator
}

// NewServer creates a new server instance
func NewServer(pdfGenerator *pdf.Generator) *Server {
	return &Server{
		pdfGenerator: pdfGenerator,
	}
}

// GeneratePDFFromPayloadHandler generates and returns a PDF report from student data in request payload
func (s *Server) GeneratePDFFromPayloadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse student data from request body
	var studentData student.Student
	if err := json.NewDecoder(r.Body).Decode(&studentData); err != nil {
		log.Printf("Error parsing student data: %v", err)
		http.Error(w, fmt.Sprintf("Invalid student data: %v", err), http.StatusBadRequest)
		return
	}

	// Generate PDF
	pdfBytes, err := s.pdfGenerator.GenerateStudentReport(&studentData)
	if err != nil {
		log.Printf("Error generating PDF: %v", err)
		http.Error(w, fmt.Sprintf("Failed to generate PDF: %v", err), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=student_%d_report.pdf", studentData.ID))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))

	// Write PDF to response
	w.Write(pdfBytes)
}

// HealthCheckHandler provides a health check endpoint
func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "healthy",
		"service": "go-student-service",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SetupRoutes configures all the routes for the server
func (s *Server) SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	
	// Health check
	router.HandleFunc("/health", s.HealthCheckHandler).Methods("GET")
	
	// PDF generation endpoint
	router.HandleFunc("/students/pdf", s.GeneratePDFFromPayloadHandler).Methods("POST")
	
	return router
}