# Screen Recorder Backend

This is the backend for the Screen Recorder application, built using Go and the Gin framework. It provides RESTful APIs for user management, screen recording functionalities, and more.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Environment Variables](#environment-variables)
- [Setup](#setup)
- [Running Migrations](#running-migrations)
- [Starting the Project](#starting-the-project)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)

## Features

- User registration and authentication
- Screen recording management
- FAQ management
- Newsletter subscription
- General settings management

## Technologies

- Go
- Gin
- GORM (ORM for Go)
- PostgreSQL
- Docker (optional for containerization)

## Environment Variables

Create a `.env` file in the root of the project with the following variables:

```bash
JWT_SECRET=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_HOST=
DB_PORT=
```


## Setup

1. **Clone the repository:**

   ```bash
   git clone https://naseeb88900@bitbucket.org/waleed-naeem/screen-recorder-backend.git
   cd screen-recorder-backend
   ```

2. **Install dependencies:**

   Make sure you have Go installed. You can install the necessary Go packages by running:

   ```bash
   go mod tidy
   ```

3. **Set up your PostgreSQL database:**

   Ensure that PostgreSQL is running and that you have created a database named `screenrecorder` (or the name you specified in the `.env` file).

## Running Migrations

To run the database migrations, use the following command:

```bash
go run cmd/migrate.go
```


This command will automatically create the necessary tables in your PostgreSQL database based on the defined models.

## Starting the Project

To start the project, use the following command:

```bash
fresh
```


This command will run the application and automatically reload it when changes are detected.

## API Documentation

Refer to the API documentation for details on the available endpoints and their usage.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.