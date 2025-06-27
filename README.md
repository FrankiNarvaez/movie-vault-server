# Movie Vault Server

This project is the backend server for the Movie Vault application, responsible for managing movie data, user authentication, and serving API requests. It's built with Go and provides a robust and scalable solution for your movie collection needs.

## Features

- **Movie Management**: CRUD operations for movies.
- **User Authentication**: Secure user login and registration with JWT.
- **Database Integration**: PostgreSQL database for data storage.
- **API Endpoints**: RESTful API for interacting with the movie vault.

## Technologies Used

- **Go**: The primary language for the server.
- **Gin Gonic**: A fast web framework for Go.
- **SQLX**: A database wrapper that provides a set of extensions to the standard `database/sql` library.
- **lib/pq**: PostgreSQL driver for Go.
- **JWT (golang-jwt/jwt/v5)**: For handling JSON Web Tokens for authentication.
- **Bcrypt (golang.org/x/crypto/bcrypt)**: For secure password hashing.

## Getting Started

### Prerequisites

- Go (version 1.20 or higher)
- PostgreSQL database

### Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/your-username/movie-vault-server.git
    cd movie-vault-server
    ```

2. **Install dependencies**:

    ```bash
    go mod tidy
    ```

3. **Database Setup**:
    - Ensure you have a PostgreSQL database running.
    - Update your database connection string in the application's configuration (e.g., in an environment variable or configuration file, as applicable for this project's structure). You might need to inspect the `db` or `internal` directories for configuration details.

### Running Database Migrations and Seeding

The project includes a script to manage database migrations and seed initial data.

To execute the script, from the root path of the project, run:

```bash
go run cmd/script [options]
```

#### Options

- `-m | --migrate`: Create database tables based on defined schemas.
- `-s | --seed`: Insert initial data into the database. (Requires tables to be created first).
- `-h | --help`: Show the help menu for the script.

**Example Usage**:

To create tables and then seed data:

```bash
go run cmd/script --migrate
go run cmd/script --seed
```

### Running the Server

To start the Movie Vault API server, execute the following command from the project root:

```bash
go run cmd/main.go
```

The server will typically run on `http://localhost:8080` or a port configured in your environment.
