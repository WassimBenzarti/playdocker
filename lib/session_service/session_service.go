package session_service

import (
	"fmt"
	"os"
	"path"

	"github.com/wassimbenzarti/pwd-go/api"
	"github.com/wassimbenzarti/pwd-go/configs"
	"github.com/wassimbenzarti/pwd-go/lib/storage"
)

type SessionService struct {
	storage storage.Storage
	cache   *Cache
}

type Cache struct {
	Cookie string `json:"cookie"`
}

func New() (*SessionService, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	parentFolder := path.Join(home, "/.pwd-go")
	if err != nil {
		return nil, err
	}

	// Instantiate storage
	storage, err := storage.New(parentFolder, "config")
	if err != nil {
		return nil, err
	}

	// Instantiate the cache
	cache := &Cache{}

	// Instantiate the session service
	sessionService := &SessionService{*storage, cache}

	return sessionService, nil

}

func (service SessionService) sync() error {
	return service.storage.Load(service.cache)
}

func (service SessionService) flush() error {
	return service.storage.Save(service.cache)
}

func (service SessionService) GetConfigPath() string {
	return service.storage.Path
}

func (service SessionService) GetSSHHost() string {
	sessionId := api.CreateSession(service.cache.Cookie)
	instance := api.CreateInstance(sessionId)
	return fmt.Sprintf(configs.SSH_URI, instance.ProxyHost)
}

func (service SessionService) HasCredentials() bool {
	service.sync()
	return service.cache.Cookie != ""
}

func (service SessionService) SetCredentials(cookie string) error {
	service.cache.Cookie = cookie
	return service.flush()
}
