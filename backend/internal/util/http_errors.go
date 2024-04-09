package util

import (
	"log"
	"main/backend/pkg/http_errors"
	"net/http"
)

func HandleError(w http.ResponseWriter, e error) {
	response := http_errors.ErrorResponse(e)
	http.Error(w, response.Error, response.Status)
	log.Println(e.Error())
}