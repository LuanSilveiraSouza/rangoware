package crypter

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	key, err := GenerateKey()

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(len(key), 32) {
		t.Fatal("Crypto key not properly generated")
	}

	test := "Hello World!"

	encrypted, err := Encrypt([]byte(test), key)

	if err != nil {
		t.Fatal(err)
	}

	decrypted, err := Decrypt(encrypted, key)

	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(string(decrypted), test) {
		t.Fatal("Encrypt/Decrypt not working as expected")
	}
}
