package abcc

import (
	"testing"
)

func TestTimestamp(t *testing.T) {
	if info, err := GetInstance().Timestamp(nil); err != nil {
		t.Fatal(err)
	} else if info.Timestamp == 0 {
		t.FailNow()
	}
}
