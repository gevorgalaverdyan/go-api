package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine){
	r.GET("/facilities", getAllFacilities)
	r.GET("/facility/:id", getFacilityById)
}