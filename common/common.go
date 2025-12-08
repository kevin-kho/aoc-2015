package common

import (
	"bytes"
	"os"
)

func ReadInput(filePath string) ([]byte, error) {
	var data []byte

	data, err := os.ReadFile(filePath)
	if err != nil {
		return data, err
	}

	return data, nil

}

func TrimNewLineSuffix(byteArr []byte) []byte {
	return bytes.TrimSuffix(byteArr, []byte{10})
}
