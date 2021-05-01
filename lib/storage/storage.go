package storage

import (
	"encoding/json"
	"os"
)

type Storage struct {
	Path string
}

func New(parentFolder string, path string) (*Storage, error) {
	// Create the parent folders
	err := os.MkdirAll(parentFolder, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &Storage{parentFolder + path}, nil
}

func (storage *Storage) Save(object interface{}) error {
	f, err := os.Create(storage.Path)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		return err
	}
	_, err = f.Write(bytes)
	return err
}

func (storage *Storage) Load(instance interface{}) error {
	bytes, err := os.ReadFile(storage.Path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, instance)
	return err
}
