<h1 align="center">
  <img src="https://raw.githubusercontent.com/create-go-app/cli/master/.github/images/cgapp_logo%402x.png" width="224px"/><br/>
  Let's Go Chat
</h1>

<p align="center"><a href="https://pkg.go.dev/github.com/create-go-app/cli/v3?tab=doc" target="_blank"><img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" /></a>&nbsp;<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" /></p>

## ⚡ Quick start
First of all, [download](https://golang.org/dl/) and install **Go**. Version `1.17` or higher is required.

## 📕 Packages

### • Hasher
#### Description:
Package provides the ability to hash and verify hashed passwords for aes standard. <br />
The **Advanced Encryption Standard (AES)** aka Rijndael is an encryption algorithm created in 2001 by NIST. It uses 128-bit blocks of data to encrypt and is a symmetric block cipher.

#### How to use:
First, create or generate a secret key. <font color="red">**IMPORTANT:** secret key must be 32 bytes in size.</font><br />
Also create a password. <font color="yellow">Note: string of any size.</font>
```go
secretKey := "di542eX9LzYA38xaH59MhT7Cr4v9cBsP"
pass := "random-password"
```

Create a new hasher entity using the constructor and pass the secret key into it:
```go
h := hasher.NewHasher(secretKey)
```

Now you can use hasher methods:
```go
h.HashPassword()
h.CheckHashPassword()
```

Example of hashing a password using a secret key:
```go
hash, err := h.HashPassword(pass)
if err != nil {
    panic(err.Error())
}
```

Example of comparing a password and its hash::
```go
if checkPass := hr.CheckHashPassword(pass, hashPass); checkPass {
    log.Println("password is correct")
	//
	// Some other code...
	//
}
```