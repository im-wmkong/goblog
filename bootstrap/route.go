package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/routes"
)

// 路由初始化
func SetupRoute() *mux.Router  {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	return router
}
