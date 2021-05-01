package session_service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wassimbenzarti/pwd-go/configs"
)

func TestSessionService(t *testing.T) {
	sessionService, err := New()
	if err != nil {
		panic(err)
	}
	sessionService.SetCredentials(configs.TEST_COOKIE)

	assert.True(t, sessionService.HasCredentials())
}
