package validplease

import "testing"

type sample struct {
	Email string `json:"email" vp:"required"`
	FirstName string `json:"firstName" vp:"maxLen(10),minLen(2)"`
	LastName string  `vp:"maxLen(10),minLen(2)"`
}

func TestMinLen(t *testing.T) {
	p := sample{Email:"amin@live.com", FirstName: "ab",LastName: "bb"}
	if e := ValidPlease(p, "en"); e != nil {
		t.Errorf("Some thing went wrong %v", e)
	}
}

func TestMaxLen(t *testing.T) {
	p := sample{Email:"a@b.com",FirstName: "Amin",LastName: "Bagheri"}
	if e := ValidPlease(p, "en"); e != nil {
		t.Errorf("Some thing went wrong %v", e)
	}
}