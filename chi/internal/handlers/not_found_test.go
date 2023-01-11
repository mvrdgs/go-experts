package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NotFound(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, err := http.NewRequest("GET", "/2", nil)
	require.NoError(t, err)

	response := ExecuteRequest(req, s)

	assert.Equal(t, "route does not exist", response.Body.String())
}
