# Simple Go HTTP Server

This is a basic HTTP server written in Go with two endpoints and request logging.

## Features

1. **Request Logging**: Logs URL, method, and response time for every request.
2. **Health Check Endpoint**: Returns a simple message to check server status.
3. **Echo Endpoint**: Returns text provided in the request.

## Endpoints

### Health Check
- **URL**: `GET /healthcheck`
- **Response**: 
  ```
  Hello World
  ```

### Echo
- **URL**: `POST /api/echo/{text}`
- **Response**: Echoes the `{text}` from the URL.

## How to Run

1. Install Go (version 1.16 or later).
2. Run the server:
   ```bash
   go run main.go
   ```
3. Access the endpoints:
   - Health Check: [http://localhost:8080/healthcheck](http://localhost:8080/healthcheck)
   - Echo: Send a POST request to `http://localhost:8080/api/echo/{text}`

## Example

### Health Check
```bash
curl -X GET http://localhost:8080/healthcheck
```

### Echo
```bash
curl -X POST http://localhost:8080/api/echo/hello
```

## Notes
- The server runs on port `8080`.
- Timeout settings: Read/Write (10s), Idle (1m).

