package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func generateBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generatePassword(s int64) (string, error) {
	b, err := generateBytes(s)
	return base64.URLEncoding.EncodeToString(b)[0:s], err
}

func main() {
	rand.Seed(time.Now().Unix())
	LENGTH := 20
	password, err := generatePassword(int64(LENGTH))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(password)
	password, err = generatePassword(int64(LENGTH))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(password)
}
