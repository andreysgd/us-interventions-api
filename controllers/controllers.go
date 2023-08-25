package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/us-interventions-api/database"
	"github.com/us-interventions-api/models"
)

func ShowInterventions(ctx *gin.Context) {
	var interventions []models.Intervention
	database.DB.Find(&interventions)
	ctx.JSON(http.StatusOK, interventions) // can switch 'http.StatusOK' to simply ' 200'
}

func InsertIntervention(ctx *gin.Context) {
	var intervention models.Intervention

	if err := ctx.ShouldBindJSON(&intervention); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error()})
		return
	}

	database.DB.Create(&intervention)
	ctx.JSON(http.StatusOK, intervention)
}

func SearchInterventionByID(ctx *gin.Context) {
	var intervention models.Intervention
	id := ctx.Params.ByName("id")
	database.DB.First(&intervention, id)

	if intervention.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": "Intervention not found."})
		return
	}

	ctx.JSON(http.StatusOK, intervention)
}

func DeleteIntervention(ctx *gin.Context) {
	var intervention models.Intervention
	id := ctx.Params.ByName("id")
	database.DB.Delete(&intervention, id)

	ctx.JSON(http.StatusOK, gin.H{
		"Sucess": "Intervention deleted."})
}

func EditIntervention(ctx *gin.Context) {
	var intervention models.Intervention
	id := ctx.Params.ByName("id")
	database.DB.First(&intervention, id)

	if err := ctx.ShouldBindJSON(&intervention); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	database.DB.Model(&intervention).UpdateColumns(intervention)
	ctx.JSON(http.StatusOK, intervention)
}

func SearchInterventionByYear(ctx *gin.Context) {
	var intervention models.Intervention
	year := ctx.Param("year")
	database.DB.Where(&models.Intervention{Year: year}).First(&intervention)

	if intervention.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": "Intervention not found."})
		return
	}

	ctx.JSON(http.StatusOK, intervention)
}

/* func SearchInterventionByCountry(ctx *gin.Context) {
	var intervention models.Intervention
	country := ctx.Param("country")
	database.DB.Where(&models.Intervention{Country: country}).First(&intervention)

	if intervention.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": "Intervention not found."})
		return
	}

	ctx.JSON(http.StatusOK, intervention)
} */
