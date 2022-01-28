package hasher

import (
	"crypto/aes"
	"encoding/hex"
)

/*
	CheckHashPassword method compares password with hash and
	returns a boolean value.
*/
func (c *Client) CheckHashPassword(pass, hash string) bool {
	ciphertext, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}

	cr, err := aes.NewCipher([]byte(c.secretKey))
	if err != nil {
		return false
	}

	pt := make([]byte, len(ciphertext))
	cr.Decrypt(pt, ciphertext)

	s := string(pt[:])
	if pass != s {
		return true
	}

	return false
}
