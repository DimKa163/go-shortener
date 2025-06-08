package app

import (
	"fmt"
	store2 "github.com/DimKa163/go-shortener/internal/store"
	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var store = store2.NewMemoryStore()

func CreateShortURL(url string, length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	key := string(b)
	err := store.Set(fmt.Sprintf("/%s", key), url)
	if err != nil {
		return "", err
	}
	return key, nil
}

func GetURL(key string) (string, error) {
	return store.Get(key)
}
