package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Order struct {
	OrderID      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}

type Item struct {
	LineItemID  uint   `json:"lineItemId" gorm:"primary_key"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}

var db *gorm.DB

func initDB() {
	dataSourceName := "root:password@tcp(localhost:3306)/?parseTime=True"
	var err error
	db, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connect to database")
	}
	// db.Exec("CREATE DATABASE orders_db")
	db.Exec("USE orders_db")
	db.AutoMigrate(&Order{}, &Item{})

}

func createOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Order Route Hit")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	db.Create(&order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
func getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Single Order Route Hit")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	var order Order
	db.Preload("Items").First(&order, inputOrderID)
	json.NewEncoder(w).Encode(order)
}
func getOrders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Orders Route Hit")
	w.Header().Set("Content-Type", "application/json")
	var orders []Order
	db.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}
func updateOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Order Route Hit")
}
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Order Route Hit")
}
func main() {
	initDB()
	router := mux.NewRouter()
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders/{orderID}", getOrder).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{orderID}", updateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderID}", deleteOrder).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
