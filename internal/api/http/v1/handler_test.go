package v1

import (
	"testing"

	"github.com/kokhno-nikolay/lets-go-chat/internal/service"
	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	h := NewHandler(&service.Services{})
	require.IsType(t, &Handler{}, h)
}
