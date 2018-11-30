package abcc

import (
	"encoding/json"
	"testing"

	"github.com/hexoul/go-abcc/types"
)

func TestMe(t *testing.T) {
	if ret, err := GetInstance().Me(nil); err != nil {
		t.Fatal(err)
	} else if ret.Email == "" {
		t.FailNow()
	} else if ret.Accounts == nil || len(ret.Accounts) == 0 {
		t.FailNow()
	}
}

func TestTrades(t *testing.T) {
	if ret, err := GetInstance().Trades(&types.Options{
		MarketCode: "metaeth",
		PerPage:    "100",
	}); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range ret.Trades {
			if b, bErr := json.Marshal(v); bErr == nil {
				t.Log(string(b))
			}
		}
	}
}
