package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Order represents the model for an order
type Order struct {
	OrderID string `json:"oderId"`
	CustomerName string `json:"customerName"`
	OrderedAt time.Time `json:"orderedAt"`
	Items []Item `json:"items"`
}

// Item represents the model for an item in the order
type Item struct {
	ItemID string `json:"itemID"`
	Description string `json:"description"`
	Quantity int `json:"quantity"`
}

var orders []Order
var prevOrderID = 0

func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/orders", createOrder).Methods(http.MethodPost)
	// Read
	router.HandleFunc("/orders/{orderId}", getOrder).Methods(http.MethodGet)
	// Read-all
	router.HandleFunc("/orders", getOrders).Methods(http.MethodGet)
	// Update
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods(http.MethodPut)
	// Delete
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application./json")
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for _, order := range orders {
		if order.OrderID == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}


