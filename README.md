# URL Shortener

A simple URL shortener service written in Go using the Gin web framework. This service allows you to shorten URLs and then redirect to the original URLs using the shortened links. It also provides metrics for the top domains.

## Features

- Shorten URLs and generate short links.
- Redirect to the original URL using the short link.
- Retrieve metrics for the top domains.

## Prerequisites

- Go 1.18 or higher

## Installation

1. **Clone the repository**

   ```sh
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener

2. **Initialize the Go module**

    ```sh
    go mod init url-shortener
    
3. **Install dependencies**

    ```sh
    go mod init url-shortener

## Usage

1. **Run the application**

    ```sh
    go run main.go
    ```

2. **Endpoints**

    - **POST /shorten**

      Shorten a new URL.

      ```sh
      curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "http://example.com"}'
      ```

      Response:

      ```json
      {
          "short_url": "http://localhost:8080/shortURL"
      }
      ```

    - **GET /:shortURL**

      Redirect to the original URL.

      ```sh
      curl -L http://localhost:8080/shortURL
      ```

    - **GET /metrics**

      Retrieve the top domains.

      ```sh
      curl http://localhost:8080/metrics
      ```

## Project Structure

    ```plaintext
    url-shortener/
    ├── handlers
    │   └── handlers.go
    ├── storage
    │   └── storage.go
    ├── utils
    │   └── utils.go
    └── main.go
      
- **main.go:** The entry point of the application. Sets up the routes and starts the server.
- **handlers/handlers.go:** Contains the HTTP handlers for shortening URLs, redirecting, and retrieving metrics.
- **storage/storage.go:** Provides in-memory storage for the URLs and metrics.
- **utils/utils.go:** Utility functions for generating short URLs.


