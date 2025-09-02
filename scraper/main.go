package main

import (
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/scraper/database"
	"github.com/MahmoodAhmed-SE/degree-progress-tracker/scraper/internals"
)

func main() {
	database.InitDB()
	internals.ScrapeMajors()
}
