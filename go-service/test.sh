#!/bin/bash

# Test script for Go Student Service
# This demonstrates how to use the service endpoints

echo "=== Go Student Service Test Script ==="
echo

# 1. Test health check
echo "1. Testing health endpoint..."
curl -s http://localhost:8080/health | jq '.' 2>/dev/null || curl -s http://localhost:8080/health
echo -e "\n"

# 2. Test with mock data (since we don't have the actual student API running)
echo "2. Testing student endpoints..."
echo "Note: These will fail with the external API since it's not running locally"
echo "But the service structure is ready for your actual student API"

echo -e "\nTesting student JSON endpoint:"
curl -s http://localhost:8080/students/2 2>/dev/null || echo "Expected: External API not available"

echo -e "\nTesting PDF generation endpoint:"
curl -s -o student_report.pdf http://localhost:8080/students/2/pdf 2>/dev/null || echo "Expected: External API not available"

# 3. Show available endpoints
echo -e "\n3. Available endpoints:"
echo "   GET /health                 - Service health check"
echo "   GET /students/{id}          - Get student details as JSON" 
echo "   GET /students/{id}/pdf      - Generate PDF report for student"

echo -e "\n4. Configuration:"
echo "   PORT: ${PORT:-8080}"
echo "   STUDENT_API_URL: ${STUDENT_API_URL:-http://localhost:3000}"

echo -e "\n5. Sample usage with your student API:"
echo "   # Set your API URL"
echo "   export STUDENT_API_URL=https://your-api.com"
echo ""
echo "   # Get student details"
echo "   curl http://localhost:8080/students/2"
echo ""
echo "   # Download PDF report"
echo "   curl -o student_report.pdf http://localhost:8080/students/2/pdf"

echo -e "\n=== Test Complete ==="