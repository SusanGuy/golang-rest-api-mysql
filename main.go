package main

import (
	"time"

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
	Description string `json:description`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}

func main() {

}
