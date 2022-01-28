package hasher

import (
	"crypto/aes"
	"encoding/hex"
)

/*
	HashPassword returns the AES hash of the password from 32-bit encryption key.
	The Advanced Encryption Standard (AES) aka Rijndael is an encryption algorithm
	created in 2001 by NIST. It uses 128-bit blocks of data to encrypt and is a
	symmetric block cipher.
*/
func (c *Client) HashPassword(pass string) (string, error) {
	cr, err := aes.NewCipher([]byte(c.secretKey))
	if err != nil {
		return "", err
	}

	out := make([]byte, len(pass))
	cr.Encrypt(out, []byte(pass))

	return hex.EncodeToString(out), nil
}
