package types

// Options for request
type Options struct {
	// Mandatory
	AccessKey string `json:"access_key,omitempty"`
	Timestamp string `json:"tonce,omitempty"`
	// Optional
	Direction   string `json:"direction,omitempty"`
	EndID       string `json:"end_id,omitempty"`
	EndTime     string `json:"end_time,omitempty"`
	MarketCode  string `json:"market_code,omitempty"`
	Page        string `json:"page,omitempty"`
	PerPage     string `json:"per_page,omitempty"`
	Size        string `json:"size,omitempty"`
	StartDealNo string `json:"startDealNo,omitempty"`
	StartID     string `json:"start_id,omitempty"`
	StartTime   string `json:"start_time,omitempty"`
	Symbol      string `json:"symbol,omitempty"`
	UtcStart    string `json:"utcStart,omitempty"`
	UtcEnd      string `json:"utcEnd,omitempty"`
	WithTrade   string `json:"withTrade,omitempty"`
}

// OptionOrder is used for ordering
// options should be sorted alphabetically with key (a-z)
var OptionOrder = []string{
	"access_key",
	"direction",
	"end_id",
	"end_time",
	"market_code",
	"page",
	"per_page",
	"size",
	"startDealNo",
	"start_id",
	"start_time",
	"symbol",
	"tonce",
	"utcStart",
	"utcEnd",
	"withTrade",
}
