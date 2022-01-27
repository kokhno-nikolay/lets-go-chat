package main

import (
	"github.com/kokhno-nikolay/lets-go-chat/pkg/hasher"
	"log"
)

func main() {
	secretKey := "124351512"
	pass := "12341241"

	h := hasher.NewHasher(secretKey)
	hashPass, err := h.HashPassword(pass)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if checkPass := hr.CheckHashPassword(pass, hashPass); checkPass {
		log.Println("password is correct")
	}
}
