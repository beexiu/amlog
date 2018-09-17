package main

import "crypto/rand"

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func uuid(L int) string {
	const charMap = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	buf := make([]byte, L)
	_, err := rand.Read(buf)
	assert(err)
	for i := 0; i < L; i++ {
		ch := buf[i]
		buf[i] = charMap[int(ch)%62]
	}
	return string(buf)
}
