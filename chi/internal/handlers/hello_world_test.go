package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HelloWorld(t *testing.T) {
	s := CreateNewServer()
	s.MountHandlers()

	req, _ := http.NewRequest("GET", "/", nil)

	response := ExecuteRequest(req, s)

	assert.Equal(t, "hello world", response.Body.String())
}
