package router

import (
	"github.com/gorilla/mux"
	"github.com/i-m-vivek/status-checker/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.WelcomeApiHander)
	router.HandleFunc("/POST/websites", controller.PostWebsiteHandler).Methods("POST")
	router.HandleFunc("/GET/websites", controller.GetWebsiteHandler).Methods("GET")

	return router
}
