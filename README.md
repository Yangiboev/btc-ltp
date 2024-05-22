## Overview

This project is a simple HTTP server written in Go that fetches the latest trade price (LTP) for specific cryptocurrency pairs from the Kraken API. The server provides an endpoint `/api/v1/ltp` that returns the latest prices for `BTC/USD`, `BTC/CHF`, and `BTC/EUR`.

## Project Structure
- `main.go`: The main application file that sets up the HTTP server and handles requests.

## Prerequisites
To run locally, ensure you have the following prerequisites installed:

- Go programming language (version 1.22.3 or higher)
- Git
- Docker

## Setup

1. Clone this repository to your local machine:

    ```bash
    git clone git@github.com:Yangiboev/btc-ltp.git
    ```

2. Navigate to the project directory:

    ```bash
    cd btc-ltp
    ```

4. Set up your PORT for the service:
   
   Create a `.env`:

    ```plaintext
    PORT=8080
    ```

## Running the Service

To run the service, execute the following command in your terminal:

```bash
go run main.go
```

## Running the Service Test

To run the integration tests, execute the following command in your terminal:

```bash
go test -v ./...
```

## Running with Docker

Build:

```bash
docker build -t ltp-server .
```

Run:
```bash
docker run -d -p 8080:8080 --name ltp-server-container ltp-server
```
## API

To try out API: https://app.danke.uz/api/v1/ltp

