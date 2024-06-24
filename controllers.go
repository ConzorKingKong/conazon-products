package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func routeIdHelper(w http.ResponseWriter, r *http.Request) (string, int, error) {
	routeId := r.PathValue("id")

	parsedRouteId, err := strconv.Atoi(routeId)
	if err != nil {
		log.Printf("Error parsing route id: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
		return "", 0, err
	}

	return routeId, parsedRouteId, nil
}

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Status: http.StatusNotFound, Message: "invalid path" + r.URL.RequestURI(), Data: ""})
}

func Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		// get all products from db (pagination)
		conn, err := pgx.Connect(context.Background(), DatabaseURLEnv)
		if err != nil {
			log.Printf("Error connecting to database: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
			return
		}

		defer conn.Close(context.Background())

		rows, err := conn.Query(context.Background(), "select * from products.products")
		if err != nil {
			log.Printf("Error getting products: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
			return
		}

		defer rows.Close()

		var rowSlice []Product

		for rows.Next() {
			var row Product
			err = rows.Scan(&row.ID, &row.CreatedAt, &row.UpdatedAt, &row.Name, &row.Description, &row.MainImage, &row.Category, &row.Price, &row.Quantity, &row.Author)
			if err != nil {
				log.Printf("Error scanning rows: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
				return
			}
			rowSlice = append(rowSlice, row)
		}

		json.NewEncoder(w).Encode(ProductsResponse{Status: http.StatusOK, Message: "Success", Data: rowSlice})
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{Status: http.StatusMethodNotAllowed, Message: "Method Not Allowed", Data: ""})
		return
	}
}

func ProductId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, parsedRouteId, err := routeIdHelper(w, r)
	if err != nil {
		return
	}

	if r.Method == "GET" {
		conn, err := pgx.Connect(context.Background(), DatabaseURLEnv)
		if err != nil {
			log.Printf("Error connecting to database: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
			return
		}

		defer conn.Close(context.Background())

		Product := Product{}

		err = conn.QueryRow(context.Background(), "select * from products.products where id=$1", parsedRouteId).Scan(&Product.ID, &Product.CreatedAt, &Product.UpdatedAt, &Product.Name, &Product.Description, &Product.MainImage, &Product.Category, &Product.Price, &Product.Quantity, &Product.Author)
		if err != nil {
			log.Printf("Error getting product with id %d - %s", parsedRouteId, err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Response{Status: http.StatusNotFound, Message: "Product not found", Data: ""})
			return
		}

		json.NewEncoder(w).Encode(ProductResponse{Status: http.StatusOK, Message: "Success", Data: Product})
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{Status: http.StatusMethodNotAllowed, Message: "Method Not Allowed", Data: ""})
		return
	}
}
