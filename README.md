# URL Shortener Service

A simple URL shortener service built with Go, Gin, and PostgreSQL. This service allows users to generate short URLs for long links and redirect short URLs to their original destinations.

## Features

- Generate a unique short URL for any valid original URL
- Store and retrieve URL mappings in a PostgreSQL database
- Redirect short URLs to their original URLs
- RESTful API endpoints

## Project Structure

```
url_shortner/
├── main.go                # Application entry point
├── go.mod                 # Go module definition and dependencies
├── database/
│   └── db.go              # Database connection and table creation
├── models/
│   └── urls.go            # URL model and database operations
├── routes/
│   ├── routes.go          # Route registration
│   └── urls.go            # Route handlers (shorten, redirect)
```

## API Endpoints

### Create Short URL

- **POST** `/shorten`
- **Body:** `{ "original_url": "https://example.com" }`
- **Response:** `{ "response": { ...shortened url object... } }`

### Redirect to Original URL

- **GET** `/{shortCode}`
- Redirects to the original URL if the short code exists.

## Setup & Run

1. **Clone the repository:**
   ```sh
   git clone https://github.com/tirlochanarora16/url_shortner.git
   cd url_shortner
   ```
2. **Set up environment variables:**
   - Create a `.env` file with your PostgreSQL connection string:
     ```env
     CONNECTION_STRING=postgres://user:password@localhost:5432/dbname?sslmode=disable
     ```
3. **Install dependencies:**
   ```sh
   go mod tidy
   ```
4. **Run the application:**
   ```sh
   go run main.go
   ```
   The server will start on `http://localhost:3000`.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [PostgreSQL](https://www.postgresql.org/) - Database
- [shortid](https://github.com/teris-io/shortid) - Short ID generator
- [godotenv](https://github.com/joho/godotenv) - Environment variable loader

## License

MIT
