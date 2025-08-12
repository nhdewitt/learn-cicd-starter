package auth

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptyAuthHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	require.Error(t, err, "no authorization header included")
}

func TestMalformedAuthHeader(t *testing.T) {
	header := make(http.Header)
	header["Authorization"] = []string{"ApuKey"}
	_, err := GetAPIKey(header)
	require.Error(t, err, "malformed authorization header")
}

func TestSuccessfulAPIKey(t *testing.T) {
	header := make(http.Header)
	header["Authorization"] = []string{"ApiKey testing"}
	key, err := GetAPIKey(header)
	require.NoError(t, err, "")
	require.NotEqual(t, "testing", key)
}
