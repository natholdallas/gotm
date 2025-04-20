package t

import (
	"encoding/json"
	"math/rand"

	"github.com/jinzhu/copier"
)

func JSONUnMarshal[T any](bytes []byte, data *T) (*T, error) {
	err := json.Unmarshal(bytes, &data)
	return data, err
}

func JSONStringify(data any) (string, error) {
	d, err := json.MarshalIndent(data, "", "\t")
	return string(d), err
}

func Copy[I any](src *I, source any) I {
	copier.CopyWithOption(src, source, copier.Option{IgnoreEmpty: true})
	return *src
}

func RandomID(length int) string {
	var result string
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	counter := 0
	for counter < length {
		result += string(chars[rand.Intn(len(chars))])
		counter++
	}
	return result
}
