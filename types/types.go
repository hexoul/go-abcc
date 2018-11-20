// Package types of both request and response for API
package types

// Request structure
type Request struct {
	AccessKey string `json:"accesskey,omitempty"`
	Timestamp string `json:"tonce,omitempty"`
	Params    string `json:"params,omitempty"`
	Sign      string `json:"signature,omitempty"`
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
