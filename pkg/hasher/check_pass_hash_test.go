package hasher

import (
	"fmt"
	"testing"
)

func TestCheckHashPassword(t *testing.T) {
	for i, tt := range []struct {
		pass   string
		secret string
		hash   string
		res    bool
	}{
		{
			"5bf07105-6aac-4220-a5bc-c90590b38a20",
			"47H37P6dRfN66DLy5rCA3sP37xdzdXkh",
			"5853fe150561d796d8e10e301db899af0000000000000000000000000000000000000000",
			true,
		},
		{
			"5bf07105-6aac-4220-a5bc-c90590b38a20",
			"47H37P6dRfN66DLy5rCA3sP37xdzdXkh",
			"bad-hash",
			false,
		},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			hasher := NewHasher(tt.secret)

			res := hasher.CheckHashPassword(tt.pass, tt.hash)
			if res != tt.res {
				t.Errorf("want %v; got %v", tt.res, res)
			}
		})
	}
}
