# Ebuy

Ebuy is an e-commerce system written in Go that generates complex analytical reports. The system handles orders, customers, products, and transactions.

## Steps to run

- Clone the repository
- Setup Postgres database
  - Pull docker image `docker pull postgres:alpine`
  - Run container `docker run --name e-buy -p 5432:5432 -e POSTGRES_PASSWORD=your_password -d postgres:alpine`
  - Run the `database/init.sql` script to create the database schema
- Create a `.env` file in the root project directory and populate all the required variables (Take the reference for environment variables from `.env.dev` file)
- Run the project `go run main.go`

