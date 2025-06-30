package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/conzorkingkong/conazon-products/config"
	"github.com/conzorkingkong/conazon-products/types"
	authhelpers "github.com/conzorkingkong/conazon-users-and-auth/helpers"
	authtypes "github.com/conzorkingkong/conazon-users-and-auth/types"
	"github.com/jackc/pgx/v5"
)


func Products(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		// get all products from db (pagination)
		conn, err := pgx.Connect(context.Background(), config.DatabaseURLEnv)
		if err != nil {
			log.Printf("Error connecting to database: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(authtypes.Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
			return
		}

		defer conn.Close(context.Background())

		rows, err := conn.Query(context.Background(), "select * from products.products")
		if err != nil {
			log.Printf("Error getting products: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(authtypes.Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
			return
		}

		defer rows.Close()

		var rowSlice []types.Product

		for rows.Next() {
			var row types.Product
			err = rows.Scan(&row.ID, &row.CreatedAt, &row.UpdatedAt, &row.Name, &row.Description, &row.MainImage, &row.Category, &row.Price, &row.Quantity, &row.Author)
			if err != nil {
				log.Printf("Error scanning rows: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(authtypes.Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
				return
			}
			rowSlice = append(rowSlice, row)
		}

		json.NewEncoder(w).Encode(types.ProductsResponse{Status: http.StatusOK, Message: "Success", Data: rowSlice})
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(authtypes.Response{Status: http.StatusMethodNotAllowed, Message: "Method Not Allowed", Data: ""})
		return
	}
}

func ProductId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, parsedRouteId, err := authhelpers.RouteIdHelper(w, r)
	if err != nil {
		return
	}

	if r.Method == "GET" {
		conn, err := pgx.Connect(context.Background(), config.DatabaseURLEnv)
		if err != nil {
			log.Printf("Error connecting to database: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(authtypes.Response{Status: http.StatusInternalServerError, Message: "Internal Service Error", Data: ""})
			return
		}

		defer conn.Close(context.Background())

		Product := types.Product{}

		err = conn.QueryRow(context.Background(), "select * from products.products where id=$1", parsedRouteId).Scan(&Product.ID, &Product.CreatedAt, &Product.UpdatedAt, &Product.Name, &Product.Description, &Product.MainImage, &Product.Category, &Product.Price, &Product.Quantity, &Product.Author)
		if err != nil {
			log.Printf("Error getting product with id %d - %s", parsedRouteId, err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(authtypes.Response{Status: http.StatusNotFound, Message: "Product not found", Data: ""})
			return
		}

		json.NewEncoder(w).Encode(types.ProductResponse{Status: http.StatusOK, Message: "Success", Data: Product})
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(authtypes.Response{Status: http.StatusMethodNotAllowed, Message: "Method Not Allowed", Data: ""})
		return
	}
}