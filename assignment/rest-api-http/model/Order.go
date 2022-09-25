package model

import (
	"hacktiv8-learning/assignment/rest-api-http/config"
)

type Order struct {
	Id           int64  `json:"id"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
}

type Item struct {
	Id          int64  `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	OrderId     int64  `json:"order_id"`
}

func GetAllData() ([]Order, error) {
	db := config.GetDb()
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	var results = []Order{}
	for rows.Next() {
		var order Order
		// for each row, scan the result into our tag composite object
		err = rows.Scan(&order.Id, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, order)
	}
	return results, nil
}
