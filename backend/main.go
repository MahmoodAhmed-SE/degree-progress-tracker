package main

import (
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	api.Init()
}
