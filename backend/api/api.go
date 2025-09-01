package api

import (
	"fmt"
	"log"
	"os"

	"github.com/MahmoodAhmed-SE/degree-progress-tracker/api/handlers"
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/engine"
	"github.com/gin-gonic/gin"
)

func Init() {
	engine.InitDB()
	server := gin.Default()
	group := server.Group("/api")
	handlers.InitRoutes(*group)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	log.Println("API server running on port " + port)
	server.Run(fmt.Sprintf(":%v", port))
}
