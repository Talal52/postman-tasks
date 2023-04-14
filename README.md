# Key-Value Store

This is a simple key-value store API that uses the Gin web framework for Go.

## Prerequisites

To run this application, you will need to have Go installed on your machine.

## Installation

1. Clone this repository: `git clone https://github.com/username/repository.git`
2. Change into the project directory: `cd repository`
3. Install the dependencies: `go mod download`

## Usage

1. Start the server: `go run main.go`
2. Make HTTP requests to `localhost:8080` to interact with the key-value store API.

### Endpoints

#### `GET /key/:input`

Retrieve the value for a given key.

#### `DELETE /delete/:input`

Delete the value for a given key.

#### `PUT /update/:input`

Update the value for a given key.

#### `POST /store`

Add a new key-value pair.