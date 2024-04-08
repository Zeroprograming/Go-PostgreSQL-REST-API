## API README

### Introduction
This is the README file for the Go PostgreSQL REST API project. This API is built using Go programming language and PostgreSQL database.

### Setup Instructions
1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Create a `.env` file in the root directory of the project and add the following environment variables:
   ```
    # Web server
    PORT=":3000"

    # Database
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASSWORD=postgres
    DB_NAME=postgres
    DB_PORT=5432
   ```
4. Make sure you have Docker and Docker Compose installed on your machine.
5. Run `docker-compose up -d` to start the PostgreSQL and pgAdmin containers.
6. Run `go run main.go` to start the Go server.

### API Endpoints
- **GET /users**: Get all users.
- **GET /user/{id}**: Get a specific user by ID.
- **POST /user**: Create a new user.
- **DELETE /user?id={id}**: Delete a user by ID.
- **GET /tasks**: Get all tasks.
- **GET /task/{id}**: Get a specific task by ID.
- **POST /task**: Create a new task.
- **PUT /task/{id}**: Update a task by ID.
- **DELETE /task/{id}**: Delete a task by ID.

### Sample Request and Response
**GET /users**
Request:
```http
GET /users HTTP/1.1
Host: localhost:3000
```
Response:
```json
[
    {
        "ID": 1,
        "FirstName": "John",
        "LastName": "Doe",
        "Email": "john@example.com"
    },
    {
        "ID": 2,
        "FirstName": "Jane",
        "LastName": "Doe",
        "Email": "jane@example.com"
    }
]
```

### Dependencies
- Go (1.16+)
- PostgreSQL
- Docker
- Docker Compose
- Gorilla Mux
- GORM
- Logrus

### Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.