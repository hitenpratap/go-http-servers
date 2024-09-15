# Go Web Server with Modular `servers` and `middleware` Packages

This project demonstrates how to build a modular web server in Go with separate packages for server implementations (`net/http`, Gin, Echo) and reusable middleware (e.g., logging).

## Features

- **Multiple Servers**: Easily switch between `net/http`, Gin, and Echo web servers.
- **Modular Design**: Separate `servers` and `middleware` packages for maintainability and scalability.
- **Custom Middleware**: Reusable logging middleware applied across different server frameworks.

## Running the Project

Run the desired server by passing the server type as a command-line argument:

- **Run `net/http` Server:**
  ```bash
  go run main.go nethttp
  ```

- **Run `gin` Server:**
  ```bash
  go run main.go gin
  ```

- **Run `echo` Server:**
  ```bash
  go run main.go echo
  ```

- **Run `fiber` Server:**
  ```bash
  go run main.go fiber
  ```
