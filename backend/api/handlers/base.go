package handlers

import (
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/api/middleware"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

var Routes []Route
var AuthRoutes []Route

func AddRoute(method, path string, handler gin.HandlerFunc, roles ...string) {
	Routes = append(Routes, Route{
		Method:      method,
		Path:        path,
		HandlerFunc: middleware.Authenticate(handler, roles),
	})
}

func AddAuthRoute(method, path string, handler gin.HandlerFunc, roles ...string) {
	Routes = append(Routes, Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handler,
	})
}

func InitRoutes(group gin.RouterGroup) {
	InitUsersRoute()
	for _, route := range Routes {
		group.Handle(
			route.Method,
			route.Path,
			route.HandlerFunc,
		)
	}
	for _, route := range AuthRoutes {
		group.Handle(
			route.Method,
			route.Path,
			route.HandlerFunc,
		)
	}
}
