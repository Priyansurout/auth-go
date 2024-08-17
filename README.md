# auth-go

# Authentication Microservice in Go

This project is a simple authentication microservice built using Go, Gin, GORM, and MySQL (running in a Docker container). It includes user registration, login, and token-based authentication using JWT.

## Features

- **User Registration**: Allows new users to sign up by providing a username, password, and age.
- **User Login**: Users can log in using their credentials and receive a JWT token.
- **JWT Authentication**: Secure routes using JWT tokens.
- **Password Hashing**: User passwords are securely hashed before storage.

## Project Structure
.
├── cmd
│ └── main.go # Entry point for the application
├── internal
│ ├── controllers # Contains handlers for various routes
│ ├── user # Contains the User model and database initialization
│ └── routes # Routing logic for the application
├── go.mod # Go module file
├── go.sum # Go dependencies file
└── .gitignore # Git ignore file


## Setup Instructions

### Prerequisites

- Go 1.20 or higher
- Docker
- MySQL running in a Docker container

### Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Priyansurout/auth-go.git
   cd auth-go

2. **Start the MySQL Docker container:**

    ```bash
    docker run --name some-sql -v mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql

3. **Initialize the database:**

    ```bash
    docker exec -it some-sql mysql -u root -p
    # Enter your root password (my-secret-pw)
    CREATE DATABASE authdb;
    exit;

4. **Run the application:**

    ```bash
    go run cmd/main.go

### DONE!

