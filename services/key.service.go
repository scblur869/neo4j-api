package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

func CreateKeyFile(kFile string, key string) {
	f, err := os.Create(kFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(key)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetEncryptionKeyEnv() string {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes) //encode key in bytes to string

	return key
}
