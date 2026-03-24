#!/bin/bash

# Demo script for the new PDF generation endpoint with student data payload
# This script demonstrates how to use the POST /students/pdf endpoint

echo "Testing the new PDF generation endpoint with student data payload..."
echo ""

# Create mock student data
STUDENT_DATA='{
  "id": 12345,
  "name": "John Smith",
  "email": "john.smith@school.edu",
  "systemAccess": true,
  "phone": "555-0123",
  "gender": "Male",
  "dob": "2005-03-15T00:00:00.000Z",
  "class": "10th Grade",
  "section": "A",
  "roll": 25,
  "fatherName": "Robert Smith",
  "fatherPhone": "555-0124",
  "motherName": "Mary Smith",
  "motherPhone": "555-0125",
  "guardianName": "Robert Smith",
  "guardianPhone": "555-0124",
  "relationOfGuardian": "Father",
  "currentAddress": "123 Main Street, Springfield, IL 62701, United States",
  "permanentAddress": "123 Main Street, Springfield, IL 62701, United States",
  "admissionDate": "2020-08-15T00:00:00.000Z",
  "reporterName": "Principal Johnson"
}'

# Check if server is running
echo "Checking if server is running..."
if ! curl -f -s http://localhost:8080/health > /dev/null; then
    echo "Server is not running. Please start the server first with 'go run .'"
    exit 1
fi
echo "Server is running ✓"
echo ""

# Generate PDF using the new endpoint
echo "Generating PDF with student data payload..."
curl -X POST \
  http://localhost:8080/students/pdf \
  -H "Content-Type: application/json" \
  -d "$STUDENT_DATA" \
  -o "generated_student_report.pdf"

if [ $? -eq 0 ]; then
    echo "PDF generated successfully! ✓"
    echo "File saved as: generated_student_report.pdf"
    
    # Check if file was created and has content
    if [ -f "generated_student_report.pdf" ] && [ -s "generated_student_report.pdf" ]; then
        FILE_SIZE=$(stat -c%s "generated_student_report.pdf" 2>/dev/null || stat -f%z "generated_student_report.pdf" 2>/dev/null)
        echo "File size: $FILE_SIZE bytes"
        echo ""
        echo "You can open the PDF file with any PDF viewer to verify the content."
    else
        echo "Warning: PDF file was not created or is empty"
    fi
else
    echo "Failed to generate PDF ✗"
    exit 1
fi