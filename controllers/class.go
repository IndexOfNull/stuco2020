package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indexofnull/stuco2020/models"
)

func GetClassMembers(c *gin.Context) {
	id := IntParam(c, "id")
	var students []models.Student
	models.GetClassMembers(&students, uint32(id))
	c.JSON(http.StatusOK, &students)
}

func GetClass(c *gin.Context) {
	id := IntParam(c, "id")
	var class models.Class
	models.GetClass(&class, uint32(id))
	c.JSON(http.StatusOK, &class)
}
