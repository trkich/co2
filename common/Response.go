package common

import (
	"encoding/json"
	"co2/common/apimessage"
	"net/http"
)

type ResponseBody struct {
	Message        string	`json:"message"`
	Response interface{}	`json:"response"`
}

func (respBody ResponseBody) ConvertToJson() string {
	var jsonData []byte
	jsonData, responseJsonErr := json.Marshal(respBody)
	if responseJsonErr != nil {
		return ""
	}
	return string(jsonData)
}

func CreateResponse(w http.ResponseWriter, messageCode int, responseObject interface{}, statusCode int)  {
	body := ResponseBody{}
	body.Message = apimessage.StatusText(messageCode)
	body.Response = responseObject
	response , _ := json.Marshal(body)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func CreateOkResponse(w http.ResponseWriter, responseObject interface{})  {
	CreateResponse(w, apimessage.Ok, responseObject, http.StatusOK)
}

func Create404Response(w http.ResponseWriter)  {
	CreateResponse(w, apimessage.NotFound, nil, http.StatusOK)
}

func PlainResponse(w http.ResponseWriter, messageCode int, responseObject interface{})  {

	response , _ := json.Marshal(responseObject)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
