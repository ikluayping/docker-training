package controller

import (
	"encoding/json"
	"fmt"
	"go-service1/pkg/model"
	"go-service1/pkg/service"
	"net/http"
)

func Healthz(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func PostLogin(w http.ResponseWriter, request *http.Request) {
	loginBean := new(model.LoginBean)
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&loginBean)
	if err != nil {
		fmt.Printf("Cannot decode json form with error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	respObj := service.LoginToService2(loginBean)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(respObj)
}
