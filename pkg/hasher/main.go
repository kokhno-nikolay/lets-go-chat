/* Package hasher provides hashing and password hash comparison. */
package hasher

type Hasher interface {
	HashPassword(pass string) (string, error)
	CheckHashPassword(pass, hash string) bool
}

type Client struct {
	secretKey string
}

func NewHasher(secretKey string) Hasher {
	return &Client{
		secretKey: secretKey,
	}
}
