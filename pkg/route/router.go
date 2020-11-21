package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Router 路由对象
var route *mux.Router

func SetRoute(r *mux.Router)  {
	route = r
}

func GetRoute() *mux.Router {
	return route
}

func Name2URL(routeName string, pairs ...string) string {
	url, err := route.Get(routeName).URL(pairs...)
	if err != nil {
		return ""
	}
	return url.String()
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}