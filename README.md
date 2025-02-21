# Fresh Data Service

The service that keeps the data "fresh".

## How to run

### Production

1. Clone the repository (`git clone https://github.com/dsc-bot/fresh-data-service.git`)
2. Run `docker build -t fresh-data-service .`
3. Run `docker run -d -e LOG_LEVEL=info fresh-data-service`

### Development

1. Clone the repository (`git clone https://github.com/dsc-bot/fresh-data-service.git`)
2. Run `go mod download`
3. Run `go mod verify`
4. Run `go run main.go`
