# Go Location Logger

A simple Go program that handles location data sent via HTTP POST requests and inserts it into a MySQL database.

## Features

- Handles JSON payloads containing location data.
- Inserts location data into a MySQL database.
- Supports parameters `device` and `user` in HTTP requests.

## Prerequisites

Before running the program, ensure you have the following:

- Go installed: [https://golang.org/doc/install](https://golang.org/doc/install)
- Docker and Docker Compose installed: [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/go-location-logger.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-location-logger
    ```

3. Install dependencies:

    ```bash
    go get -u github.com/sirupsen/logrus
    ```

4. Create an `.env` file in the project root with your environment variables:

    ```plaintext
    # .env

    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_HOST=your_db_host
    DB_PORT=your_db_port
    DB_NAME=your_db_name
    ```

## Usage

### Running with Docker Compose

1. Create a Docker network:

    ```bash
    docker network create app-network
    ```

2. Use Docker Compose to start the services:

    ```bash
    docker-compose up -d
    ```

3. Access your Go application at `http://localhost:8080`. Make HTTP POST requests as described in the README.md.

4. To stop the services, run:

    ```bash
    docker-compose down
    ```

### Running without Docker

If you prefer not to use Docker, you can run the Go application directly using:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

Make a POST request with a JSON payload to `http://localhost:8080/?device=your_device&user=your_user`. The JSON payload should include location data.

Example:

```bash
curl -X POST -d '{"_type": "location", "tst": 1637650367, "lat": 37.7749, "lon": -122.4194, "tid": "123", "batt": 90, "vac": 220}' http://localhost:8080/?device=your_device&user=your_user
```

## Contributing

Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.