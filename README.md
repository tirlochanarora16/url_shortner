# URL Shortener Service

A simple URL shortener service built with Go, Gin, and PostgreSQL. This service allows users to generate short URLs for long links and redirect short URLs to their original destinations.

## Features

- Generate a unique short URL for any valid original URL
- Store and retrieve URL mappings in a PostgreSQL database
- Redirect short URLs to their original URLs
- **Redis-based access counting** for high-performance hit tracking
- **Bulk database updates** via Python cron job for efficient data persistence
- RESTful API endpoints
- **Basic SQL migration tool for schema changes**
- **Automated CI/CD pipeline with GitHub Actions, AWS EC2, and RDS PostgreSQL**

## Project Structure

```
url_shortner/
├── main.go                # Application entry point
├── go.mod                 # Go module definition and dependencies
├── database/
│   ├── db.go              # Database connection and table creation
│   ├── migrations.go      # Basic SQL migration tool
├── models/
│   └── urls.go            # URL model and database operations
├── routes/
│   ├── routes.go          # Route registration
│   └── urls.go            # Route handlers (shorten, redirect)
├── pkg/
│   └── redisClient.go     # Redis client configuration
├── scripts/
│   └── redis_sync.py      # Python cron job for bulk Redis to DB sync
├── .github/
│   └── workflows/
│       └── deploy.yml     # GitHub Actions workflow for CI/CD
```

## SQL Migration Tool

This project includes a **basic SQL migration tool** implemented in `database/migrations.go`. It allows you to:

- Define migrations as Go structs with a unique name and SQL query
- Apply schema changes (such as adding columns) in a transactional way
- Track applied migrations in a `schema_migrations` table to prevent duplicate execution

**Example migration struct:**

```go
Migrations{
    Table:         "urls",
    ColumnName:    "updated_at",
    Quey:          AlterUrlTableUpdateAt, // SQL query string
    MigrationName: "add_updated_at_1_06_25",
}
```

**Key methods:**

- `ApplyMigration()`: Runs the migration SQL and records it in the database
- `AddMigrationToDB()`: Adds the migration name to the `schema_migrations` table

## CI/CD Pipeline with GitHub Actions, AWS EC2, and RDS PostgreSQL

This project uses **GitHub Actions** for continuous integration and deployment. The workflow is defined in `.github/workflows/deploy.yml` and provides:

- **Automatic build and deployment** on every push to the `main` branch
- **Go binary build** for Linux (suitable for EC2 deployment)
- **Secure upload** of the built binary to an AWS EC2 instance using SSH and SCP
- **Environment variable management** by creating a `.env` file on the EC2 instance with the PostgreSQL connection string (for AWS RDS)
- **Graceful restart** of the application on the EC2 instance
- **Integration with AWS RDS PostgreSQL** for persistent database storage

**Key steps in the workflow:**

- Build the Go binary for Linux
- Upload the binary to EC2 using `appleboy/scp-action`
- SSH into EC2, set up environment variables, kill any running instance, and start the new binary with the correct environment

This setup enables seamless, automated deployments to AWS infrastructure whenever you push code to the repository.

## Redis-Based Access Counting

This project implements a **high-performance access counting system** using Redis for real-time hit tracking and a Python cron job for bulk database updates.

### How It Works

1. **Real-time Hit Tracking**: When a short URL is accessed, the system immediately increments a Redis counter using the key pattern `url_hits:{shortCode}`. This provides instant response times without database writes.

2. **Bulk Database Updates**: A Python cron job (`scripts/redis_sync.py`) runs periodically to:
   - Read all hit counters from Redis
   - Perform bulk updates to the PostgreSQL database
   - Clear processed counters from Redis

### Benefits

- **High Performance**: Redis operations are extremely fast, ensuring minimal latency during URL redirects
- **Scalability**: Can handle thousands of concurrent requests without database bottlenecks
- **Data Consistency**: Periodic bulk updates ensure data persistence while maintaining performance
- **Fault Tolerance**: Redis provides persistence and replication options for reliability

### Setup Redis

1. **Install Redis** (if not already installed):
   ```sh
   # macOS
   brew install redis
   
   # Ubuntu/Debian
   sudo apt-get install redis-server
   ```

2. **Start Redis**:
   ```sh
   redis-server
   ```

3. **Configure Redis connection** in your environment:
   ```env
   REDIS_URL=redis://localhost:6379
   ```

### Python Cron Job Setup

1. **Install Python dependencies**:
   ```sh
   pip install redis psycopg2-binary
   ```

2. **Set up the cron job** to run the sync script periodically:
   ```sh
   # Add to crontab (runs every 5 minutes)
   */5 * * * * /usr/bin/python3 /path/to/url_shortner/scripts/redis_sync.py
   ```

3. **Configure the sync script** with your database and Redis connection details.

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
   - Create a `.env` file with your PostgreSQL connection string and Redis URL:
     ```env
     CONNECTION_STRING=postgres://user:password@localhost:5432/dbname?sslmode=disable
     REDIS_URL=redis://localhost:6379
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
- [Redis](https://redis.io/) - In-memory data store for access counting
- [shortid](https://github.com/teris-io/shortid) - Short ID generator
- [godotenv](https://github.com/joho/godotenv) - Environment variable loader

### Python Dependencies (for cron job)

- `redis` - Redis client for Python
- `psycopg2-binary` - PostgreSQL adapter for Python

## License

MIT
