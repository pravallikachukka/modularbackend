# Backend README

# Overview
This directory contains the backend for Le Petite Patisserie project, that serves menu data via an API or to a front end that uses the API.

The project is written in Golang, and uses the MUX library for handling requests, and GORM for object ralational mapping to a database.

The project utilises it's own SQLite database when running locally.



# Endpoints
## GET
`/menu` - Gets all items in the menu and returns as JSON <br>
`/menu/id` - Returns a specific item from the menu for a given ID

## POST
N/A
## PUT
N/A

# Running The Server Locally
To run the server lcoally, run `go run src/main.go`

In the browser http://localhost:8080 to test

# Example request
`curl localhost:8080/menu`

# Docker Run Command
TBC

# Libraries
GORM - https://gorm.io/
Mux - https://github.com/gorilla/mux
