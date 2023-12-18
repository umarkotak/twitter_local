# TWITTER LOCAL
A sample twitter like project for learning golang

## Code structure
- Route
- Controller
- Service
- Repository

## Pre-requisites
1. Postgres database as the data store
  - make a db user, with name: admin & password: admin
  - make a new database with name: twitter_local_dev
2. Rename `.env.sample` into `.env`
  - this will act as an environment variable source

## Run the project
1. To run the app
  - `make run`

---

## Additional installation
1. Install swaggo
  - `go install github.com/swaggo/swag/cmd/swag@latest`
