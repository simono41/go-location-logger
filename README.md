# Go Location Logger

A simple Go program that handles location data sent via HTTP POST requests and inserts it into a MySQL database.

## Features

- Handles JSON payloads containing location data.
- Inserts location data into a MySQL database.
- Utilizes headers `X-Limit-U` and `X-Limit-D` for specifying user and device information in HTTP requests.

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

1. Use Docker Compose to start the services:

    ```bash
    docker-compose up -d
    ```

2. Access your Go application at `http://localhost:8080`. Make HTTP POST requests as described in the README.md.

3. Import the Grafana Dashboard from `dashboard.json`. Use the Grafana UI to import the dashboard and visualize your data.

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
curl -X POST -H "Content-Type: application/json" \
  -H "X-Limit-U: your_user" \
  -H "X-Limit-D: your_device" \
  -H "Authorization: Basic your_token" \
  -d '{
    "_type": "location",
    "tst": 1700820453,
    "lat": 37.7749,
    "lon": -122.4194,
    "alt": 90,
    "batt": 94,
    "acc": 35,
    "bs": 2,
    "p": 99.314,
    "created_at": 1700820457,
    "BSSID": "e8:48:b8:7f:b4:d4",
    "SSID": "Mr.Puhu",
    "vac": 20,
    "tag": "Arbeit",
    "topic": "owntracks/simono41/6193B679-AD67-4B93-9DF2-158501A055AF",
    "conn": "w",
    "m": 1,
    "tid": "AF"
  }' \
  http://localhost:8080/
```

## Contributing

Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.