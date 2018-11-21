// Package types of both request and response for API
package types

// Error structure
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

// Response structure
type Response struct {
	Error *Error `json:"error"`
}

// Account structure
type Account struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
	Locked   string `json:"locked"`
}

// UserInfo structure
type UserInfo struct {
	SN        string     `json:"sn"`
	Name      int64      `json:"name"`
	Email     string     `json:"email"`
	Activated bool       `json:"activated"`
	Accounts  []*Account `json:"accounts"`
}

// Timestamp structure
type Timestamp struct {
	Timestamp float64 `json:"timestamp"`
}

// Market structure
type Market struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	BaseUnit    string `json:"base_unit"`
	QuoteUnit   string `json:"quote_unit"`
	PriceFixed  int    `json:"price_fixed"`
	VolumeFixed int    `json:"volume_fixed"`
}

// Markets structure
type Markets struct {
	Markets []*Market `json:"markets"`
}

// Meta structure
type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
}

// Order structure
type Order struct {
}

// Orders structure
type Orders struct {
	Meta   *Meta    `json:"meta"`
	Orders []*Order `json:"orders"`
}
