package abcc

import "testing"

func TestOrders(t *testing.T) {
	if _, err := GetInstance().Orders(nil); err != nil {
		t.Fatal(err)
	}
}
