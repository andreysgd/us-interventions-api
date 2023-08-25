package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/us-interventions-api/controllers"
)

func HandleRequest() {
	router := gin.Default()
	router.GET("/interventions", controllers.ShowInterventions)
	router.POST("/interventions", controllers.InsertIntervention)
	router.GET("/interventions/:id", controllers.SearchInterventionByID)
	router.DELETE("/interventions/:id", controllers.DeleteIntervention)
	router.PATCH("/interventions/:id", controllers.EditIntervention)
	router.GET("/interventions/search/:year", controllers.SearchInterventionByYear)
	// router.GET("/interventions/search/:country", controllers.SearchInterventionByCountry)
	router.Run() // listen and serve on localhost:8080 (can be changed by passing parameter as ":8000")
}
