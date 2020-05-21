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
	OrderID      string    `json:"oderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}

// Item represents the model for an item in the order
type Item struct {
	ItemID      string `json:"itemID"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

var orders []Order
var prevOrderID = 0

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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

// CreateOrder godoc
// @Summary Create new order
// @Description Create new order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	order.OrderID = strconv.Itoa(prevOrderID)
	orders = append(orders, order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// GetOrder godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application./json")
	json.NewEncoder(w).Encode(orders)
}

// GetOrder godoc
// @Summary Get details of a specific order
// @Description Get details of a specific order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} Order
// @Router /orders/orderId [get]
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

// UpdateOrder godoc
// @Summary Update a specific order
// @Description Update a specific order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} Order
// @Router /orders/orderId [put]
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

// DeleteOrder godoc
// @Summary Delete a specific order
// @Description Delete a specific order
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {object} Order
// @Router /orders/orderId [delete]
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
