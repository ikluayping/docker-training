package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-service1/pkg/model"
	"net"
	"net/http"
	"os"
	"time"
)

func LoginToService2(bean *model.LoginBean) model.ResponseLogin {
	transport := http.Transport{
		Dial: dialTimeout,
	}
	client := http.Client{
		Transport: &transport,
	}

	byteBuffer := new(bytes.Buffer)
	_ = json.NewEncoder(byteBuffer).Encode(bean)

	service2name := os.Getenv("SERVICE_2_NAME")

	req, err := http.NewRequest("POST", "http://"+service2name+":8100/login", byteBuffer)
	if err != nil {
		fmt.Printf("Cannot new request %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Cannot login with service2 module %s", err.Error())
		return model.ResponseLogin{
			ResponseCode: "10",
			ResponseDesc: err.Error(),
		}
	}

	if resp.StatusCode == http.StatusOK {
		return model.ResponseLogin{
			ResponseCode: "00",
			ResponseDesc: "login success",
		}
	}

	return model.ResponseLogin{
		ResponseCode: "10",
		ResponseDesc: "Login failed",
	}
}

func dialTimeout(network, addr string) (net.Conn, error) {
	timeout := time.Duration(1000 * time.Millisecond)
	return net.DialTimeout(network, addr, timeout)
}
