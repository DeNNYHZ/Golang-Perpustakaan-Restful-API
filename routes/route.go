package routes

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/controller"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/seed", migrate.MigrateKeDB).Methods("GET", "OPTIONS")
	router.HandleFunc("/register", middlewares.RenderKeJSON(controller.Register)).Methods("POST", "OPTIONS")
	router.HandleFunc("/loginadmin", middlewares.RenderKeJSON(controller.LoginAdmin)).Methods("POST")
	router.HandleFunc("/api/home", middlewares.RenderKeJSON(middlewares.HarusAuth(controller.Home))).Methods("GET")
	router.HandleFunc("/api/refresh", middlewares.RenderKeJSON(middlewares.HarusAuth(controller.RefreshToken))).Methods("POST")

	return router
}
