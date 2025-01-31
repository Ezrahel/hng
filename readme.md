# Go API - User Details with CORS Support

This is a simple Go API that provides user details in JSON format. The API includes basic CORS (Cross-Origin Resource Sharing) support, allowing it to be accessed from different domains. It dynamically generates the current date and time in ISO 8601 (UTC) format.

## Features

- Returns user details including:
  - Email address
  - Current date and time in ISO 8601 format (UTC)
  - GitHub link
- Handles CORS (Cross-Origin Resource Sharing), allowing it to be used across different domains.
- Supports pre-flight requests (`OPTIONS` method) for cross-origin HTTP requests.
- Implements dynamic date generation for the `current_datetime` field.

## API Endpoint

### `GET /user`

- **Description**: Returns the user details in JSON format with the current UTC date and time in ISO 8601 format.
- **Response**: JSON object containing the user's email, GitHub link, and the current date-time.

### Example Response:

```json
{
  "email": "adelakinisrael024@gmail.com",
  "current_datetime": "2025-01-30T15:04:05Z",
  "github_link": "https://github.com/ezrahel/hng-one"
}
```

### CORS Headers:

- `Access-Control-Allow-Origin`: Allows requests from any origin (set to `*` for this example, but should be restricted in production).
- `Access-Control-Allow-Methods`: Specifies allowed methods for cross-origin requests (GET, POST).
- `Access-Control-Allow-Headers`: Allows custom headers like `Content-Type`.

## Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.15 or higher).
- A terminal/command prompt to run the API.

## Installation and Running the API

1. Clone this repository:

   ```bash
   git clone https://github.com/yourusername/go-api-user-details.git
   cd go-api-user-details
   ```

2. Run the Go application:

   ```bash
   go run main.go
   ```

3. The API will start and listen on port `8080`. You should see the following message:

   ```
   Server is listening on port 8000...
   ```

4. To test the API, open your browser or use `curl`:

   ```bash
   curl http://localhost:8000
   ```

   Example response:

   ```json
   {
     "email": "adelakinisrael024@gmail.com",
     "current_datetime": "2025-01-30T15:04:05Z",
     "github_link": "https://github.com/ezrahel/hng"
   }
   ```

## Handling CORS

This API includes basic CORS support:

- The `Access-Control-Allow-Origin` header allows the API to accept requests from any origin (currently set to `*`).


If you are making requests from a frontend application on a different domain, CORS headers will allow the browser to successfully make the request.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to replace any placeholders (like the GitHub repository URL) with actual values. This `README` should help users understand how to install, run, and use the API, along with the key features and functionality.

Let me know if you need any changes!