# Conazon Products

This is the products endpoint for the Conazon project.

## Quickstart

To test locally, setup a `.env` file in the root directory with the following variables:

`DATABASEURL` - Url to postgres database. REQUIRED

Datbase url should bne formatted this - 'host=localhost port=5432 user=postgres password={password} dbname={dbname} sslmode=disable'

Then run:

`go build .`
`./conazon-products`

## Endpoints (later will have swagger)

- /

GET - generic hello world. useless endpoint

- /products

GET - Returns all products

- /products/{id}

GET - Returns a single product
