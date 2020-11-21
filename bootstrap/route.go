package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/pkg/route"
	"goblog/routes"
)

// 路由初始化
func SetupRoute() *mux.Router  {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	return router
}
