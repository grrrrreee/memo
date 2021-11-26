package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// https://m.blog.naver.com/pjt3591oo/221354809830
// 위 링크에서 받아옴. 조금은 변경시켜야 할 듯
func hashing() {
	// 입력. 정보를 받아오는 것으로 변경
	data := "123"
	hash := sha256.New()

	hash.Write([]byte(data))

	result := hash.Sum(nil)
	resultst := hex.EncodeToString(result)

	// 출력.
	fmt.Println(resultst)
}