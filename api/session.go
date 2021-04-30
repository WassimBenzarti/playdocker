package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/wassimbenzarti/pwd-go/configs"
	"github.com/wassimbenzarti/pwd-go/models"
)


func CreateSession(cookie string) string {
    request, err := http.NewRequest("POST", configs.API_URL, nil)
    if err !=nil {
        log.Printf("Couldn't create the request %v", err)
    }
    request.Header.Add("cookie", cookie)

    response, err := http.DefaultClient.Do(request)
    if err != nil{
        log.Printf("Cannot request the API %v", err)
    }

    url := strings.Split(response.Request.URL.String(),"/")
    id := url[len(url)-1]
    fmt.Printf("id: %s", id)

    return id
}

func GetSession(id string) models.Session {
    sessionUrl := fmt.Sprintf("%s/sessions/%s", configs.API_URL, id)
    response, err := http.Get(
        sessionUrl,
    )
    if err != nil {
        log.Printf("Cannot get the session from URL: %s", sessionUrl)
    }
    body, err := ioutil.ReadAll(response.Body)
    if err !=nil {
        log.Println("Cannot read the body of the response")
    }
    session := models.Session{}

    if err:= json.Unmarshal(body, &session); err!=nil   {
        log.Println("Cannot parse the json returned")
    }

	return session
}