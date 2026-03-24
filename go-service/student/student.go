package student

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Student represents the student data structure
type Student struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	SystemAccess       bool   `json:"systemAccess"`
	Phone              string `json:"phone"`
	Gender             string `json:"gender"`
	DOB                string `json:"dob"`
	Class              string `json:"class"`
	Section            string `json:"section"`
	Roll               int    `json:"roll"`
	FatherName         string `json:"fatherName"`
	FatherPhone        string `json:"fatherPhone"`
	MotherName         string `json:"motherName"`
	MotherPhone        string `json:"motherPhone"`
	GuardianName       string `json:"guardianName"`
	GuardianPhone      string `json:"guardianPhone"`
	RelationOfGuardian string `json:"relationOfGuardian"`
	CurrentAddress     string `json:"currentAddress"`
	PermanentAddress   string `json:"permanentAddress"`
	AdmissionDate      string `json:"admissionDate"`
	ReporterName       string `json:"reporterName"`
}

// Service handles student-related operations
type Service struct {
	baseURL string
	client  *http.Client
}

// NewService creates a new student service
func NewService(baseURL string) *Service {
	return &Service{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetStudent fetches student details by ID
func (s *Service) GetStudent(id string) (*Student, error) {
	url := fmt.Sprintf("%s/api/v1/students/%s", s.baseURL, id)
	
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch student: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var student Student
	if err := json.NewDecoder(resp.Body).Decode(&student); err != nil {
		return nil, fmt.Errorf("failed to decode student data: %w", err)
	}

	return &student, nil
}