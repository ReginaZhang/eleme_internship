package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	debug := false
	key := os.Args[1]
	infile, err := os.Open(os.Args[2])
	plaintext, err := ioutil.ReadAll(infile)

	ciphertext, err := encrypt([]byte(key), plaintext)
	if err != nil {
		logrus.Fatal(err)
	}

	ioutil.WriteFile(os.Args[3], ciphertext, 0600)

	if debug {
		fmt.Println("before encryption: ", string(plaintext))
		fmt.Printf("after encryption: %0x\n", string(ciphertext))

		decrypted, err := decrypt([]byte(key), ciphertext)
		if err != nil {
			logrus.Error("failed to decrypt")
		}

		fmt.Println("after decryption: ", string(decrypted))
	}
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))
	return ciphertext, nil
}

func decrypt(key, encoded []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(encoded) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := encoded[:aes.BlockSize]
	encoded = encoded[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(encoded, encoded)
	return encoded, nil
}
