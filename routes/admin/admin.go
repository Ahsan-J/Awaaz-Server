package admin

import ("github.com/gorilla/mux")

func GenerateAdminRoutes (router *mux.Router){
	admin := router.PathPrefix("/admin").Subrouter()

	GenerateUserRoutes(admin)
}