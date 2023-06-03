package routes

import (
	"github.com/blackviking27/go-data-management/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterDataStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/data/", controllers.CreateData).Methods("POST")
	router.HandleFunc("/data/", controllers.GetData).Methods("GET")
	router.HandleFunc("/data/{dataId}", controllers.GetDataById).Methods("GET")
	router.HandleFunc("/data/{dataId}", controllers.UpdateData).Methods("PUT")
	router.HandleFunc("/data/{dataId}", controllers.DeleteData).Methods("DELETE")
}
