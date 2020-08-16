package gcd

// Euclid ...
func Euclid(a, b uint64) uint64 {
	if a == b {
		return a
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return Euclid(b, a%b)
}
