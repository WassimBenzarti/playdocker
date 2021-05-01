package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wassimbenzarti/pwd-go/configs"
	"github.com/wassimbenzarti/pwd-go/lib"
	"github.com/wassimbenzarti/pwd-go/models"
)

func CreateInstance(sessionId string) models.Instance {
	url := fmt.Sprintf("%s/sessions/%s/instances", configs.API_URL, sessionId)
	instance := models.Instance{}
	response, err := http.Post(url, "application/json", nil)
	err = lib.ToJson(&instance, response, err)
	if err != nil {
		log.Println("Cannot get the instance from the API")
	}
	return instance
}
