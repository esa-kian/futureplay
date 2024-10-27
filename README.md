# Matchmaking Service

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Configuration](#configuration)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Directory Structure](#directory-structure)
- [License](#license)
- [Note](#note)

## Description

The Matchmaking Service is a Golang-based API designed for grouping players into competitions based on their level and geographical location. Players can join competitions, and the service will match them with others of similar skill levels in a given timeframe.

## Features

- Players can join matchmaking queues based on their country and level.
- Automatic grouping of players into competitions of 10 based on their geographical location and level.
- Players are notified of their competition assignments once enough players have joined or after a timeout.
- Lightweight and efficient, built using Go and Docker.

## Prerequisites

- [Go 1.20+](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)

## Configuration

The application settings, including server port, competition size, and wait time, can be customized in the `config.yaml` file:

```yaml
server:
  port: 8080

matchmaking:
  competition_size: 10
  wait_time_seconds: 30
```
Adjust these values as needed before starting the application.


## Installation

### Clone the Repository

```bash
git clone https://github.com/esa-kian/futureplay.git
cd futureplay
```

### Build the Application
If you want to run the application locally without Docker, execute the following commands:

``` bash
go mod tidy
go build -o futureplay ./cmd/server.go
```

## Running the Application
### Using Docker
1. Build the Docker Image:

```bash
docker build -t futureplay .
```
2. Run the Docker Container:
```bash
docker run -p 8080:8080 futureplay
```
### Without Docker
After building the application, run it using:

```bash
./futureplay
```
The application will be accessible at `http://localhost:8080`.

## API Endpoints
### Join Matchmaking
- URL: `/matchmaking/join`

- Method: `POST`

- Request Body:
```json
{
    "id": "1",
    "level": 5,
    "country": "US"
}
```
- Response:
    - `200 OK` if successfully added to the queue.
    - `400 Bad Request` if the input is invalid.

### Example Request
Using `curl` to join matchmaking:
```bash
curl -X POST http://localhost:8080/matchmaking/join -H "Content-Type: application/json" -d '{"id": "1", "level": 5, "country": "US"}'
```

## Testing
To run unit tests, execute:
```bash
go test ./internal/storage -v # for testing storage implementation
go test ./internal/api -v # for testing HTTP request handlers
go test ./internal/service -v # for testing logic layer
```
## Directory Structure
```plaintext
futureplay/
│
├── cmd/
│   └── server.go       # Entry point of the application
│
├── internal/
│   ├── api/
│   │   └── handler.go # HTTP request handlers
│   │
│   ├── config/
│   │   └── config.go   # configuration loader 
│   │
│   ├── model/
│   │   └── model.go   # Models definition
│   │
│   ├── service/
│   │   └── matchmaking.go # Matchmaking logic
│   │
│   └── storage/
│       └── memory.go   # In-memory storage implementation
│
├── config.yaml          # configurations
├── Dockerfile           # Dockerfile for containerization
├── go.mod               # Go module file
└── README.md            # Project documentation
```

## License
This project is open-source and available under the MIT License.


## Note:
The unit tests and README documentation were generated with the assistance of ChatGPT to ensure comprehensive coverage and clear explanations.
