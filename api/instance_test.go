package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wassimbenzarti/pwd-go/configs"
)

func TestCreateInstance(t *testing.T) {
	sessionId := CreateSession(configs.TEST_COOKIE)
	instance := CreateInstance(sessionId)

	assert.NotEmpty(t, instance.Name)

	newSession := GetSession(sessionId)
	newInstance := newSession.Instances[instance.Name]
	assert.NotEmpty(t, newInstance)
	assert.NotEmpty(t, newInstance.ProxyHost)
}
