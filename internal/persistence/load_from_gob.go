package persistence

import (
	"bytes"
	"encoding/gob"
	"os"
)

func LoadFromGob[T any](filename string) ([]T, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	buffer.Write(data)
	decoder := gob.NewDecoder(&buffer)

	var result []T
	err = decoder.Decode(&result)
	return result, err
}
