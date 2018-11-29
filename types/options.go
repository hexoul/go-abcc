package types

// Options for request
type Options struct {
	AccessKey   string `json:"access_key,omitempty"`
	MarketCode  string `json:"market_code,omitempty"`
	Symbol      string `json:"symbol,omitempty"`
	Timestamp   string `json:"tonce,omitempty"`
	UtcStart    string `json:"utcStart,omitempty"`
	UtcEnd      string `json:"utcEnd,omitempty"`
	StartDealNo string `json:"startDealNo,omitempty"`
	Size        string `json:"size,omitempty"`
}

// OptionOrder is used for ordering
// options should be sorted alphabetically with key (a-z)
var OptionOrder = []string{
	"access_key",
	"market_code",
	"size",
	"startDealNo",
	"symbol",
	"tonce",
	"utcStart",
	"utcEnd",
	"withTrade",
}
