package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hashing() {
	data := "123"
	hash := sha256.New()

	hash.Write([]byte(data))

	result := hash.Sum(nil)
	resultst := hex.EncodeToString(result)

	fmt.Println(resultst)
}