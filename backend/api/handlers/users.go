package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUsersRoute() {
	AddRoute(http.MethodGet, "/v1/users", func(ctx *gin.Context) {

	}, "ADMIN")
}
