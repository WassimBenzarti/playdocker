package storage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestingInstance struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func TestStorage(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	parentFolder := home + "/.pwd-go/"
	os.MkdirAll(parentFolder, os.ModePerm)
	storage := &Storage{parentFolder + "/config"}
	instance := &TestingInstance{"Random name", 99}
	storage.Save(instance)

	newInstance := &TestingInstance{}
	storage.Load(&newInstance)

	assert.Equal(t, newInstance.Name, instance.Name)
	assert.Equal(t, newInstance.Number, instance.Number)

	defer func() {
		err := os.RemoveAll(parentFolder)
		if err != nil {
			panic(err)
		}
	}()
}
