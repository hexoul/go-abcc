package abcc

import (
	"testing"
)

func TestMe(t *testing.T) {
	if info, err := GetInstance().Me(nil); err != nil {
		t.Fatal(err)
	} else if info.Email == "" {
		t.FailNow()
	} else if info.Accounts == nil || len(info.Accounts) == 0 {
		t.FailNow()
	}
}
