# Conazon Products

This is the products endpoint for the Conazon project.

## Quickstart

To test locally, setup a `.env` file in the root directory with the following variables:

```
DATABASEURL - Url to postgres database. REQUIRED
PORT - Port to run server on. Defaults to 8081
```

Datbase url should be formatted like this if using `docker-compose` - `'host=postgres port=5432 user=postgres dbname=conazon sslmode=disable'`

Then run:

`docker-compose up`

## Endpoints (later will have swagger)

- /

GET - Catch all 404

- /products

GET - Returns all products

- /products/{id}

GET - Returns a single product
