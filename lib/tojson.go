package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ToJson(object interface{}, response *http.Response, err error) error {
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	json.Unmarshal(body, &object)
	if err != nil {
		return err
	}
	return nil
}
