package integrationtests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	response := struct {
		Success string `json:"success"`
	}{}
	requestHelper(
		t,
		"POST",
		"/write",
		strings.NewReader(`{"message":"text"}`),
		&response,
	)
	require.Equal(t, "ok", response.Success)

	var result []struct {
		ID      int64  `json:"id"`
		Message string `json:"text"`
	}
	requestHelper(t, "GET", "/read", nil, &result)

	require.Equal(t, 1, len(result))
	require.Equal(t, "text", result[0].Message)
}
