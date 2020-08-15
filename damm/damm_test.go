package damm

import (
	"testing"
)

func TestDamm(t *testing.T) {
	tt := map[string]bool{
		"1234567890": false,
		"1234567894": true,
		"2345678945": true,
		"3456789452": true,
		"4567894520": false,
	}
	for s, e := range tt {
		c := int8(s[len(s)-1] - '0')
		r := Calculate(s[0 : len(s)-1])
		if (c == r) != e {
			t.Errorf("%s: expected %v, got %v", s, c, r)
			return
		}
		g := Validate(s)
		if g != e {
			t.Errorf("%s: expected %v, got %v", s, e, g)
			return
		}
	}
}
