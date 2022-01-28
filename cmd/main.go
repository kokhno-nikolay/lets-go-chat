package main

import (
	"fmt"

	"github.com/kokhno-nikolay/lets-go-chat/pkg/hasher"
)

func main() {
	/* Create or generate a secret key and password */
	secretKey := "47H37P6dRfN66DLy5rCA3sP37xdzdXkh"
	pass := "some-random-password"

	fmt.Printf("PASSWORD: %s\n", pass)
	fmt.Printf("SECRET KEY: %s\n", secretKey)

	/* Create a new hasher entity using the constructor and pass the secret key */
	h := hasher.NewHasher(secretKey)

	fmt.Println("HASHING PASSWORD...")

	/* Pass password to hash method */
	hash, err := h.HashPassword(pass)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("HASH: %s\n", hash)

	fmt.Println("CHECKING HASH PASSWORD...")

	/* Check if the password hash matches the hash */
	if checkPass := h.CheckHashPassword(pass, hash); checkPass {
		fmt.Println("SUCCESS! PASSWORD IS CORRECT.")
	} else {
		fmt.Println("FAILED! PASSWORD AND HASH DO NOT MATCH.")
	}
}
