# Student Report Generation System

## Architecture Overview

This system implements a microservices architecture for student report generation, where the Node.js backend acts as a gateway to the Go service for PDF generation.

## System Flow

```
Frontend (React) → Node.js API → Go Service → PDF Generation
```

1. **Frontend**: Student details page with a "Download Report" button
2. **Node.js API**: Handles authentication and data processing (`/api/v1/students/:id/report`)
3. **Go Service**: Stateless PDF generation service (`POST /students/pdf`)
4. **Response**: PDF file delivered directly to the user for download

## Why This Architecture?

### 1. **Single Responsibility Principle**
- **Node.js Service**: Handles authentication, data validation, and business logic
- **Go Service**: Focuses solely on PDF generation with optimal performance

### 2. **Security Benefits**
- Authentication is centralized in the Node.js API
- Go service remains stateless and doesn't handle user credentials
- Reduced attack surface - only one endpoint needs comprehensive security measures

### 3. **Scalability**
- Go service can be horizontally scaled with multiple replicas
- Stateless design allows for easy load balancing
- CPU-intensive PDF generation is isolated from the main application

### 4. **Performance**
- Go's superior performance for CPU-intensive PDF generation
- Concurrent processing capabilities for multiple report requests
- Optimized memory usage for large document generation

## Technical Implementation

### Backend API Endpoint
```typescript
GET /api/v1/students/:id/report
```

- Validates student existence
- Fetches complete student data from database
- Calls Go service with student payload
- Returns PDF as binary stream

### Go Service Endpoint
```go
POST /students/pdf
Content-Type: application/json
```

**Request Payload:**
```json
{
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
  "currentAddress": "123 Main Street, Springfield, IL 62701",
  "permanentAddress": "123 Main Street, Springfield, IL 62701",
  "admissionDate": "2020-08-15T00:00:00.000Z",
  "reporterName": "Principal Johnson"
}
```

### Frontend Integration
- **RTK Query**: Integrated with existing API architecture
- **Loading States**: Shows "Generating Report..." during processing
- **Error Handling**: User-friendly error messages
- **Auto Download**: Automatic PDF download with proper filename

## File Structure
```
backend/
├── src/modules/students/
│   ├── students-service.js     # PDF report generation logic
│   ├── students-controller.js  # Route handler
│   └── sudents-router.js      # Route definition

frontend/
├── src/domains/student/
│   ├── api/student-api.ts     # RTK Query endpoints
│   └── components/views/mini-avatar.tsx  # Download button component

go-service/
├── handlers/handlers.go       # HTTP handlers
├── student/student.go        # Student data structures
└── pdf/pdf.go               # PDF generation logic
```

## Sample Output

A sample generated PDF (`generated_student_report.pdf`) is included in the repository for reference.

## Future Enhancements

- **Caching**: Implement Redis caching for frequently requested reports
- **Batch Processing**: Support for generating multiple student reports
- **Templates**: Configurable PDF templates for different report types
- **Analytics**: Track report generation metrics
- **Testing**: Comprehensive unit and integration tests

## Security Considerations

- JWT token validation in Node.js API
- Input validation and sanitization
- Rate limiting for PDF generation requests
- CORS configuration for cross-origin requests
- Error handling without exposing sensitive information