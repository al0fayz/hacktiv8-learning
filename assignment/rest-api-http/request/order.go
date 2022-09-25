package request

import "time"

type OrderCreate struct {
	CustomerName string       `json:"customerName"`
	OrderedAt    time.Time    `json:"orderedAt"`
	Items        []ItemCreate `json:"items"`
}

type ItemCreate struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
}

// update
type OrderUpdate struct {
	OrderId      int64        `json:"orderId"`
	CustomerName string       `json:"customerName"`
	OrderedAt    time.Time    `json:"orderedAt"`
	Items        []ItemUpdate `json:"items"`
}

type ItemUpdate struct {
	ItemId      int64  `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
}
