# Retail Pulse Image Processing Service

A service to process store images and calculate their perimeters.

## Description

This service provides APIs to:

1. Submit jobs with image URLs and store IDs.
2. Process images by calculating their perimeters.
3. Check job status.

## Assumptions

1. Images are accessible via HTTP/HTTPS URLs.
2. Images are in common formats (JPEG, PNG).
3. All store IDs in requests are valid.
4. In-memory storage is sufficient for this demo (in production, would use a database).

### Image Parameter Storage

Although the service processes images and calculates their perimeters, the image data itself is not stored permanently. Only the job status and related results (like perimeters) are stored, and the response returned to the user does not include the images themselves.

## Setup Instructions

### Without Docker

1. Install Go 1.21 or later.
2. Clone the repository.
3. Run:
   ```bash
   go mod download
   go mod tidy    
   go run cmd/server/main.go
   ```

### With Docker

1. Build the image:

   ```bash
   docker build -t retail-pulse .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 retail-pulse
   ```

## Testing

### Submit a Job

```bash
curl -X POST http://localhost:8080/api/submit \
  -H "Content-Type: application/json" \
  -d '{
    "count": 2,
    "visits": [
      {
        "store_id": "S00339218",
        "image_url": [
          "https://www.gstatic.com/webp/gallery/2.jpg",
          "https://www.gstatic.com/webp/gallery/3.jpg"
        ],
        "visit_time": "2023-12-20T10:00:00Z"
      },
      {
        "store_id": "S01408764",
        "image_url": [
          "https://www.gstatic.com/webp/gallery/3.jpg"
        ],
        "visit_time": "2023-12-20T11:00:00Z"
      }
    ]
  }'
```

### Check Job Status

If you're using zsh (Z shell), make sure to escape special characters like `?`, `&`, and `=` to prevent globbing (filename expansion).

For zsh, use one of the following methods:

#### Option 1: Escape the special characters with backslashes (`\`)

```bash
curl http://localhost:8080/api/status\?jobid=<job_id>
```

#### Option 2: Quote the entire URL to prevent interpretation of special characters

```bash
curl "http://localhost:8080/api/status?jobid=<job_id>"
```

Replace `<job_id>` with the actual job ID you received from the `SubmitJob` endpoint.

## Development Environment

- Go 1.21
- Visual Studio Code with Go extension
- Libraries:
  - gorilla/mux: HTTP router
  - google/uuid: UUID generation

## Future Improvements

1. Add persistent storage (e.g., PostgreSQL).
2. Implement a job queue system (e.g., Redis Queue).
3. Add authentication.
4. Implement proper logging and monitoring.
5. Implement proper error handling and recovery.
6. Add API documentation using Swagger/PostMan.
7. Add support for Kubernetes deployment for scalability.
