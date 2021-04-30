package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wassimbenzarti/pwd-go/configs"
	"github.com/wassimbenzarti/pwd-go/models"
)

func CreateInstance(sessionId string) models.Instance {
	url := fmt.Sprintf("%s/%s/instances", configs.API_URL, sessionId)
	response, err := http.Post(url)
	if err != nil {
		log.Println("Cannot get the instance from the API")
	}
	
}
