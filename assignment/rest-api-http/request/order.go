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
