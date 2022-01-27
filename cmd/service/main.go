package main

import (
	"fmt"

	"github.com/kokhno-nikolay/lets-go-chat/pkg/hasher"
)

func main() {
	secretKey := "47H37P6dRfN66DLy5rCA3sP37xdzdXkh"
	pass := "some-random-password"

	fmt.Printf("PASSWORD: %s\n", pass)
	fmt.Printf("SECRET KEY: %s\n", secretKey)

	fmt.Println("HASHING PASSWORD...")
	h := hasher.NewHasher(secretKey)
	hash, err := h.HashPassword(pass)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("HASH: %s\n", hash)

	fmt.Println("CHECKING HASH PASSWORD...")
	if checkPass := h.CheckHashPassword(pass, hash); checkPass {
		fmt.Println("SUCCESS! PASSWORD IS CORRECT.")
	} else {
		fmt.Println("FAILED! PASSWORD AND HASH DO NOT MATCH.")
	}
}
