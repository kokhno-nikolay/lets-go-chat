package hasher

import (
	"crypto/aes"
	"encoding/hex"
)

func (c *Client) HashPassword(pass string) (string, error) {
	cr, err := aes.NewCipher([]byte(c.secretKey))
	if err != nil {
		return "", err
	}

	out := make([]byte, len(pass))
	cr.Encrypt(out, []byte(pass))

	return hex.EncodeToString(out), nil
}
