package main

import (
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/admin-api/database"
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/admin-api/router"
)

func main() {
	database.InitDB()
	router.InitApi()
}
