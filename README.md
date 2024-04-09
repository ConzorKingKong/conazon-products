# Conazon Products

This is the products endpoint for the Conazon project.

## Quickstart

To test locally, setup a `.env` file in the root directory with the following variables:

`DATABASEURL` - Url to postgres database. REQUIRED

Datbase url should be formatted like this if using `docker-compose up` - 'host=postgres port=5432 user=postgres dbname=conazon sslmode=disable'

Then run:

`docker-compose up`

## Endpoints (later will have swagger)

- /

GET - generic hello world. useless endpoint

- /products

GET - Returns all products

- /products/{id}

GET - Returns a single product
