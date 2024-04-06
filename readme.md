# Boilerplate Golang

This boilerplate provides the basic structure for building Golang applications by:

```
* REST API
* JWT authentication
* Middleware
* Logging
```

# Structure

Project structure:

```
app
├── http
│   ├── controllers
│       ├── v1
│           ├── user_controller.go
│   ├── middlewares
│       ├── auth_middleware.go
│       ├── request_logger.go
├── modules
│   ├── auth
│       ├── service.go
│   ├── jwtgenerator
│       ├── jwtgenerator.go
│   ├── user
│       ├── entity.go
│       ├── formatter.go
│       ├── input.go
│       ├── repository.go
│       ├── service.go
config
├── config.go
databases
├── 1_create_users_table.sql
helper
├── api_response.go
├── errors_log.go
├── env.go
routes
├── routes.go
├── v1
    ├── routes_v1.go
storage
├── logs
    ├── api
    ├── errors
.env.example
```

# Steps

```
1. Clone this boilerplate to your local directory.
2. Create a `.env` file in the project root directory and fill it with the database configuration and JWT.
3. Run `go mod download` to install dependencies.
4. Run `go run main.go` to run the application.
```

# References

```
Golang: https://go.dev/
JWT: https://jwt.io/
```

# Contribution

```
This boilerplate is still under development. You can contribute by:

* File a bug report
* Submit a pull request
* Provide suggestions and ideas
```