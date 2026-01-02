# simx-go-todo

A simple TODO REST API built with Go, Gin, and PostgreSQL, featuring JWT authentication and environment-based configuration.

---

## Features

- **CRUD API** for TODO items
- **JWT Authentication** middleware
- **PostgreSQL** database integration
- **Auto-migration** for the `todos` table
- **Environment-based** configuration (no hardcoded credentials)
- **Structured project layout** for scalability

---

## Project Structure

```
simx-go-todo/
├── cmd/api/main.go           # Application entry point
├── internal/
│   ├── config/               # Database configuration (InitDB, DB pool)
│   ├── share/                # Shared middleware (logging, JWT)
│   └── todo/                 # Todo domain: repository, usecase, routes
├── go.mod
├── go.sum
└── .env                      # Environment variables (not committed)
```

---

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt)
- [jackc/pgx](https://github.com/jackc/pgx)
- [joho/godotenv](https://github.com/joho/godotenv)

### Environment Variables

Create a `.env` file in the project root:

```
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
JWT_SECRET=
PORT=8080
```

**Never commit your `.env` file or credentials to source control.**

---
## Running the Application

### Without Docker

1. Start PostgreSQL and ensure your database is ready.
2. Set up your `.env` file as described above.
3. Run the application:

   ```sh
   go run cmd/api/main.go
   ```

   Or build and run:

   ```sh
   go build -o simx-go-todo ./cmd/api
   ./simx-go-todo
   ```

   The server will start on `http://localhost:8080`.

---

### With Docker

1. Ensure you have Docker installed.
2. Create a `.env` file as described above.
3. Create a `Dockerfile` in the project root (if not present):

   ```dockerfile
   FROM golang:1.20-alpine

   WORKDIR /app

   COPY go.mod go.sum ./
   RUN go mod download

   COPY . .

   RUN go build -o simx-go-todo ./cmd/api

   EXPOSE 8080

   CMD ["./simx-go-todo"]
   ```

4. (Optional) Create a `docker-compose.yml` for both app and PostgreSQL:

   ```yaml
   version: '3.8'
   services:
     db:
       image: postgres:15
       environment:
         POSTGRES_USER: user
         POSTGRES_PASSWORD: password
         POSTGRES_DB: tododb
       ports:
         - "5432:5432"
       volumes:
         - pgdata:/var/lib/postgresql/data
     app:
       build: .
       env_file: .env
       ports:
         - "8080:8080"
       depends_on:
         - db
   volumes:
     pgdata:
   ```

5. Build and run with Docker Compose:

   ```sh
   docker-compose up --build
   ```

   Or build and run manually:

   ```sh
   docker build -t simx-go-todo .
   docker run --env-file .env -p 8080:8080 simx-go-todo
   ```

---

## API Endpoints

| Method | Endpoint      | Description         | Auth Required |
|--------|--------------|---------------------|--------------|
| GET    | /todos       | List all todos      | Yes          |
| POST   | /todos       | Create a new todo   | Yes          |
| PUT    | /todos/:id   | Update a todo       | Yes          |
| DELETE | /todos/:id   | Delete a todo       | Yes          |

**All endpoints require a valid JWT in the `Authorization: Bearer <token>` header.**

---

## Middleware

- **Global Logging:** Logs each request and response.
- **JWT Authentication:** Validates JWT tokens using the secret from `JWT_SECRET`.

---

## Database Migration

On startup, the app auto-creates the `todos` table if it does not exist.

---

## Extending

- Add more fields to the `Todo` struct as needed.
- Implement user management and token issuance for a full authentication flow.

---

## License

MIT