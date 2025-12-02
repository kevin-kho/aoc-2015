package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	input := "bgvyzdsv"
	MAX := 999999999

	for i := range MAX {

		suffix := string(i)

		whole := fmt.Sprintf("%v%v", input, suffix)

		res := md5.Sum([]byte(whole))

		if bytes.Equal(res[:5], []byte{0, 0, 0, 0, 0}) {
			fmt.Println(whole)
			fmt.Println("---")
			fmt.Println(res)
			break
		}

	}

	res := md5.Sum([]byte(input))

	fmt.Println(hex.EncodeToString(res[:]))

}
