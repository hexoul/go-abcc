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
	SN        int64      `json:"sn"`
	Name      int64      `json:"name"`
	Email     string     `json:"email"`
	Activated bool       `json:"activated"`
	Accounts  []*Account `json:"accounts"`
}
