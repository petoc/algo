package damm64

import (
	"testing"
)

type st struct {
	r bool
	e error
}

func TestDamm64(t *testing.T) {
	tt := []struct {
		b string
		r map[string]st
		e error
	}{
		{
			b: "12",
			r: map[string]st{},
			e: ErrBaseOutOfRange,
		},
		{
			b: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@$%",
			r: map[string]st{},
			e: ErrBaseOutOfRange,
		},
		{
			b: "0123456789",
			r: map[string]st{
				"1234567890": st{false, nil},
				"1234567894": st{true, nil},
				"2345678945": st{true, nil},
				"3456789452": st{true, nil},
				"4567894520": st{false, nil},
			},
			e: nil,
		},
		{
			b: "0123456789ABCDEFGHIJKLMNOPQRSTUV",
			r: map[string]st{
				"12344": st{true, nil},
				"G12QF": st{true, nil},
				"W567M": st{false, ErrInvalidCharacter},
			},
			e: nil,
		},
		{
			b: "456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			r: map[string]st{
				"12347": st{false, ErrInvalidCharacter},
				"G12QB": st{false, ErrInvalidCharacter},
				"W567O": st{true, nil},
			},
			e: nil,
		},
		{
			b: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz@$",
			r: map[string]st{
				"12344": st{true, nil},
				"G#2QB": st{false, ErrInvalidCharacter},
				"1AfF$": st{true, nil},
				"fNhF%": st{false, ErrInvalidCharacter},
				"W567M": st{false, nil},
			},
			e: nil,
		},
	}
	for _, tc := range tt {
		if len(tc.r) == 0 {
			_, err := Calculate(tc.b, "")
			if err != tc.e {
				t.Errorf("%v: expected %v, got %v", tc.b, tc.e, err)
				return
			}
			continue
		}
		for s, e := range tc.r {
			c := string(s[len(s)-1])
			r, err := Calculate(tc.b, s[0:len(s)-1])
			if err != e.e && (c == r) != e.r {
				t.Errorf("%v: %s: expected %v, got %v", tc.b, s, c, r)
				return
			}
			g := Validate(tc.b, s)
			if g != e.r {
				t.Errorf("%v: %s: expected %v, got %v", tc.b, s, e, g)
				return
			}
		}
	}
}
