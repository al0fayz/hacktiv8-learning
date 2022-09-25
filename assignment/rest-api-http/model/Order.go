package model

import (
	"hacktiv8-learning/assignment/rest-api-http/config"
	"hacktiv8-learning/assignment/rest-api-http/request"
	"hacktiv8-learning/assignment/rest-api-http/utils"
)

type Order struct {
	Id           int64  `json:"id"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
	Items        []Item `json:"items"`
}

type Item struct {
	Id          int64  `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	OrderId     int64  `json:"order_id"`
}

// get all data
func GetAllData() ([]Order, error) {
	db := config.GetDb()
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results = []Order{}
	for rows.Next() {
		var order Order
		// for each row, scan the result into our tag composite object
		err = rows.Scan(&order.Id, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			return nil, err
		}
		items, err := GetAllItems(order.Id)
		if err != nil {
			return nil, err
		}
		order.Items = items
		results = append(results, order)
	}
	return results, nil
}
func GetAllItems(orderId int64) ([]Item, error) {
	db := config.GetDb()
	rows, err := db.Query("SELECT * FROM items WHERE order_id = ?", orderId)
	if err != nil {
		return nil, err
	}
	var items = []Item{}
	defer rows.Close()
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.Id, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// create
func CreateOrder(order *request.OrderCreate) error {
	db := config.GetDb()
	sqlQuery := `
		INSERT INTO orders (customer_name, ordered_at) 
		VALUES (?, ?)
		`
	res, err := db.Exec(sqlQuery, order.CustomerName, order.OrderedAt)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	//save items
	for _, item := range order.Items {
		sqlItems := `INSERT INTO items (item_code, description, quantity, order_id) VALUES (?, ?, ?, ?)`
		_, err := db.Exec(sqlItems, item.ItemCode, item.Description, item.Quantity, id)
		if err != nil {
			return err
		}
	}
	return nil
}

// detail
func DetailOrder(id string) (*Order, error) {
	db := config.GetDb()
	sqlStatement := `SELECT * FROM orders WHERE id = ?`
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var order Order
	if rows.Next() {
		// for each row, scan the result into our tag composite object
		err = rows.Scan(&order.Id, &order.CustomerName, &order.OrderedAt)
		if err != nil {
			return nil, err
		}
		items, err := GetAllItems(order.Id)
		if err != nil {
			return nil, err
		}
		order.Items = items
	} else {
		return nil, utils.ShowError("Order not found")
	}
	return &order, nil
}

// update
func UpdateOrder(order *request.OrderUpdate) error {
	db := config.GetDb()
	sqlQuery := `
		UPDATE orders SET customer_name = ?, ordered_at = ?
		WHERE id = ?
		`
	_, err := db.Exec(sqlQuery, order.CustomerName, order.OrderedAt, order.OrderId)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	//save items
	for _, item := range order.Items {
		sqlItems := `UPDATE items SET item_code = ?, description = ?, quantity = ? WHERE id = ? AND order_id = ?`
		_, err := db.Exec(sqlItems, item.ItemCode, item.Description, item.Quantity, item.ItemId, order.OrderId)
		if err != nil {
			return err
		}
	}
	return nil
}

// delete
func DeleteOrder(id string) error {
	db := config.GetDb()
	sqlQuery := `DELETE FROM orders WHERE id = ?`
	_, err := db.Exec(sqlQuery, id)
	if err != nil {
		return err
	}
	//delete items
	sqlItems := `DELETE FROM items WHERE order_id = ?`
	_, err = db.Exec(sqlItems, id)
	if err != nil {
		return err
	}
	return nil
}
