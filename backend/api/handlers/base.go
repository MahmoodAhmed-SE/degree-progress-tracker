package handlers

import "github.com/gin-gonic/gin"

type Route struct {
	Method      string
	Path        string
	Name        string
	HandlerFunc gin.HandlerFunc
}

var Routes []Route

func AddRoute(name, method, path string, handler gin.HandlerFunc, roles ...string) {
	Routes = append(Routes, Route{
		Method:      method,
		Path:        path,
		Name:        name,
		HandlerFunc: handler,
	})
}

func InitRoutes(group gin.RouterGroup) {
	for _, route := range Routes {
		group.Handle(
			route.Method,
			route.Path,
			route.HandlerFunc,
		)
	}
}
