package approuter

import (
	"github.com/gorilla/mux"
	"go-service1/pkg/controller"
)

func NewCreateAppRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/devops/login", controller.PostLogin).Methods("OPTIONS", "POST")
	router.HandleFunc("/healthz", controller.Healthz).Methods("OPTIONS", "GET")
	return router
}
