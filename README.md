# Go JSON API Starter

[![Go Report Card](https://goreportcard.com/badge/github.com/mounis-bhat/go-rest-starter)](https://goreportcard.com/report/github.com/mounis-bhat/go-rest-starter)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Introduction

This is a simple REST API written in Go. It uses the go standard library for the most part, to keep things simple. For routing, it uses gorilla/mux. The database used is PostgreSQL. The API is secured using JWT.

## Features

- JWT authentication
- CRUD operations

## Getting Started

### Prerequisites

To run this project, you need to have Go installed on your machine. Make sure you also have a PostgreSQL database available.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/mounis-bhat/go-rest-starter.git
   cd go-rest-starter
   ```

   ```bash
   Copy the .env.example file to .env and configure the database connection details.
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Build the application:

   ```bash
    make build
   ```

4. Run the application:

   ```bash
   make run
   ```

5. Visit http://localhost:8080 to access the API.

6. To run the tests:

   ```bash
   make test
   ```

## Usage

Use the following endpoints to access the API:

| Endpoint | Method | Description            |
| -------- | ------ | ---------------------- |
| /account | GET    | Get account details    |
| /account | PATCH  | Update account details |
| /account | POST   | Create a new account   |

## License

This project is licensed under the MIT License.
