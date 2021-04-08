package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	key, err := GenerateKey()

	if err != nil {
		panic(err)
	}

	hexKey := hex.EncodeToString(key)

	directory, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(directory+"/key.txt", []byte(hexKey), 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Generated key at key.txt")
}

func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := io.ReadFull(rand.Reader, key)

	if err != nil {
		return nil, err
	}

	return key, nil
}
