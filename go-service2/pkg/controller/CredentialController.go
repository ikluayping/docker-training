package controller

import (
	"encoding/json"
	"fmt"
	"go-devops-service2/pkg/model"
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
	respObj := &model.ResponseLogin{
		ResponseCode: "00",
		ResponseDesc: "Login Success",
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(respObj)
}
