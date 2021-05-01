package api

import (
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/wassimbenzarti/pwd-go/configs"
	"github.com/wassimbenzarti/pwd-go/models"
)

var id = ""

func TestSession(t *testing.T) {
	id := CreateSession(configs.TEST_COOKIE)

	IsType(t, "", id)
}

func TestGetSession(t *testing.T) {
	id := CreateSession(configs.TEST_COOKIE)
	session := GetSession(id)
	IsType(t, models.Session{}, session)
	NotEmpty(t, session.Host)
	NotEmpty(t, session.Id)
	NotEmpty(t, session.PwdIPAddress)
	Equal(t, true, session.Ready)
}
