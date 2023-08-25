package main

import (
	"github.com/us-interventions-api/database"
	"github.com/us-interventions-api/routes"
)

func main() {
	database.DatabaseConnection()
	routes.HandleRequest()
}