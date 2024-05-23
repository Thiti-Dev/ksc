package persistence

import (
	"bytes"
	"encoding/gob"
	"os"
)

func SaveToGob[T any](filename string, data []T) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, buffer.Bytes(), 0644)
}
