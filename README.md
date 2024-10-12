# Goku Consumer

This project implements an MQTT consumer for processing URLs and uploading batches to an API.

## Getting Started

1. Configure the application in `configs/config.json`.
2. Run `go mod tidy` to download dependencies.
3. Build the application: `go build ./cmd/goku-consumer`
4. Run the application: `./goku-consumer`

## Project Structure

- `cmd/goku-consumer/`: Main application entry point
- `internal/`: Internal packages
- `pkg/`: Shared packages
- `configs/`: Configuration files

## License

[Add your chosen license here]
