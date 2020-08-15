package luhn

import (
	"errors"
	"log"
	"strconv"
	"testing"
)

type st struct {
	r bool
	e error
}

func TestLuhn(t *testing.T) {
	tt := map[string]st{
		"1234567897": st{true, nil},
		"2345678912": st{true, nil},
		"3456789123": st{false, nil},
		"4567891234": st{false, nil},
		"56789123xx": st{false, strconv.ErrSyntax},
	}
	for s, e := range tt {
		c, err := Calculate(string(s[:len(s)-1]))
		log.Println(err)
		r := string(s[len(s)-1])
		if err != e.e && !errors.Is(err, e.e) {
			t.Errorf("%s: expected [%v], got [%v]", s, e.e, err)
			return
		}
		if (strconv.FormatInt(int64(c), 10) == r) != e.r {
			t.Errorf("%s: expected %v, got %v", s, r, c)
			return
		}
		g := Validate(s)
		if g != e.r {
			t.Errorf("%s: expected %v, got %v", s, e.r, g)
			return
		}
	}
}
