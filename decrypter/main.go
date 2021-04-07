package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Decrypter \nInsert decrypt key:")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)

	fmt.Println(text)

	os.Exit(3)
}

func Decrypt(cypherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	plainText, err := gcm.Open(nil, cypherText[:gcm.NonceSize()], cypherText[gcm.NonceSize():], nil)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}
