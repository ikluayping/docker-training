package controller

import (
	"encoding/json"
	"fmt"
	"go-service2/pkg/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Healthz use for checking livness and readiness probe
func Healthz(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//PostLogin is a function for validate Login from request
func PostLogin(w http.ResponseWriter, request *http.Request) {
	respObj := &model.ResponseLogin{
		ResponseCode: "00",
		ResponseDesc: "Login Success",
	}
	loginBean := new(model.LoginBean)
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&loginBean)
	if err != nil {
		fmt.Printf("Cannot decode json form with error %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonvalid, err := os.Open("valid-user.json")
	if err != nil {
		log.Printf("Cannot open file name valid-user.json with error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		respObj.ResponseCode = "10"
		respObj.ResponseDesc = err.Error()
		_ = json.NewEncoder(w).Encode(respObj)
		return
	}

	defer jsonvalid.Close()
	byteValue, _ := ioutil.ReadAll(jsonvalid)
	var result map[string]interface{}
	_ = json.Unmarshal([]byte(byteValue), &result)

	if loginBean.Username != result["username"] || loginBean.Password != result["password"] {
		w.WriteHeader(http.StatusBadRequest)
		respObj.ResponseCode = "10"
		respObj.ResponseDesc = "Username or Password is invalid."
		_ = json.NewEncoder(w).Encode(respObj)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(respObj)
}
