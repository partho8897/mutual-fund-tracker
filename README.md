# Mutual Fund Tracker API

A RESTful API in Go for tracking Indian mutual funds. The API allows users to list all mutual funds, search for specific funds, retrieve historical and latest data for a fund, and backtest mutual fund data. It uses the Gin framework for routing and provides structured error handling.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installing](#installing)
- [API Endpoints](#api-endpoints)
    - [List All Mutual Funds](#list-all-mutual-funds)
    - [Search for a Mutual Fund](#search-for-a-mutual-fund)
    - [Get Historical Data for a Mutual Fund](#get-historical-data-for-a-mutual-fund)
    - [Get Latest Data for a Mutual Fund](#get-latest-data-for-a-mutual-fund)
    - [Backtest Mutual Fund Data](#backtest-mutual-fund-data)
- [Usage](#usage)
- [Contributing](#contributing)

## Features

- List all mutual funds
- Search for mutual funds by name
- Retrieve historical data for a specific mutual fund
- Retrieve the latest data for a specific mutual fund
- Backtest mutual fund data with user-provided inputs

## Getting Started

Follow these instructions to set up and run the project on your local machine for development and testing purposes.

### Prerequisites

- Go 1.22 or higher

### Installing

1. Clone the repository:

    ```sh
    git clone https://github.com/partho8897/mutual-fund-tracker.git
    cd mutual-fund-tracker
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Run the application:

    ```sh
    go run main.go
    ```

## API Endpoints

### List All Mutual Funds

- **URL:** `/mftracker/v1/list`
- **Method:** `GET`
- **Description:** Retrieves a list of all mutual funds.

### Search for a Mutual Fund

- **URL:** `/mftracker/v1/search`
- **Method:** `GET`
- **Query Parameters:**
    - `fundName` (required): The name of the mutual fund to search for.
- **Description:** Searches for mutual funds by name.

### Get Historical Data for a Mutual Fund

- **URL:** `/mftracker/v1/fund/:fundId`
- **Method:** `GET`
- **Path Parameters:**
    - `fundId` (required): The ID of the mutual fund.
- **Description:** Retrieves historical data for a specific mutual fund.

### Get Latest Data for a Mutual Fund

- **URL:** `/mftracker/v1/fund/:fundId/latest`
- **Method:** `GET`
- **Path Parameters:**
    - `fundId` (required): The ID of the mutual fund.
- **Description:** Retrieves the latest data for a specific mutual fund.

### Backtest Mutual Fund Data

- **URL:** `/mftracker/v1/backtrack`
- **Method:** `POST`
- **Request Body:** JSON object containing backtest parameters.
- **Description:** Performs a backtest on mutual fund data based on user-provided inputs.

## Usage

To interact with the API, you can use tools like `curl`, Postman, or any HTTP client of your choice. Below are some examples:

- **List all mutual funds:**
    ```sh
    curl -X GET http://localhost:80/mftracker/v1/list
    ```

- **Search for a mutual fund:**
    ```sh
    curl -X GET "http://localhost:80/mftracker/v1/search?fundName=HDFC"
    ```

- **Get historical data for a mutual fund:**
    ```sh
    curl -X GET http://localhost:80/mftracker/v1/fund/101281
    ```

- **Get the latest data for a mutual fund:**
    ```sh
    curl -X GET http://localhost:80/mftracker/v1/fund/101281/latest
    ```

- **Backtest mutual fund data:**
    ```sh
    curl --location 'http://localhost:80/mftracker/v1/backtrack' \
    --header 'Content-Type: application/json' \
    --data '{
      "from": "01-01-2022",
      "investmentType": "Lumpsum",
      "investmentFrequency": "Monthly",
      "fundInfos": [
        {
          "SchemeCode": "101281",
          "amount": 10000
        },
        {
          "SchemeCode": "100822",
          "amount": 10000
        }
      ]
    }'
    ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your improvements.