package api

import (
	"os"

	"github.com/MahmoodAhmed-SE/degree-progress-tracker/api/handlers"
	"github.com/gin-gonic/gin"
)

func Init() {
	server := gin.Default()
	group := server.Group("/api")
	handlers.InitRoutes(*group)
	port := os.Getenv("PORT")
	server.Run(port)
}
