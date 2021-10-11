package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Order struct {
	idOrder   string     `json:"idorder"`
	Customer  string     `json:"customer"`
	Inventory *Inventory `json:"inventory"`
}

type Inventory struct {
	idInven string `json:"idinven"`
	Name    string `json:"name"`
	Qty     string `json:"qty"`
}

var orders []Order
var inventories []Inventory

func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func getInventories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(inventories)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.idOrder = strconv.Itoa(rand.Intn(1000))
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

func main() {
	r := mux.NewRouter()

	orders = append(orders, Order{idOrder: "1", Customer: "Cust 1", Inventory: &Inventory{idInven: "1", Name: "Hijab", Qty: "5"}})
	orders = append(orders, Order{idOrder: "2", Customer: "Cust 2", Inventory: &Inventory{idInven: "2", Name: "Gamis", Qty: "-2"}})

	inventories = append(inventories, Inventory{idInven: "1", Name: "Hijab", Qty: "5"})
	inventories = append(inventories, Inventory{idInven: "2", Name: "Gamis", Qty: "-2"})

	r.HandleFunc("/orders", getOrders).Methods("GET")
	r.HandleFunc("/inventories", getInventories).Methods("GET")
	r.HandleFunc("/orders", createOrder).Methods("POST")

	fmt.Printf("Server start at port :8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
