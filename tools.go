package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+-"

type Tools struct{}

func (t *Tools) RandomString(n int) string {
	var (
		s, r = make([]rune, n), []rune(randomStringSource)
		l    = len(randomStringSource)
		y    = uint64(l)
		x    uint64
	)
	for i := range s {
		prime, _ := rand.Prime(rand.Reader, l)
		x = prime.Uint64()
		s[i] = r[x%y]
	}
	return string(s)
}
