package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kokhno-nikolay/lets-go-chat/internal/config"
	"github.com/kokhno-nikolay/lets-go-chat/internal/service"
)

func TestNewHandler(t *testing.T) {
	h := NewHandler(&service.Services{})
	require.IsType(t, &Handler{}, h)
}

func TestHandler_Init(t *testing.T) {
	h := NewHandler(&service.Services{})

	router := h.Init(config.GetConfig())
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}

	require.Equal(t, http.StatusOK, res.StatusCode)
}
