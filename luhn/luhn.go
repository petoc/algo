package luhn

import "strconv"

// Calculate ...
func Calculate(s string) (int, error) {
	l := len(s)
	c, err := strconv.Atoi(string(s[l-1]))
	if err != nil {
		return 0, err
	}
	p := l % 2
	for i, r := range s[0 : l-1] {
		n := int(r - '0')
		if i%2 == p {
			n = n * 2
		}
		if n > 9 {
			n = n - 9
		}
		c = c + n
	}
	return c % 10, nil
}

// Validate ...
func Validate(s string) bool {
	c, err := Calculate(s)
	return err == nil && c == 0
}
