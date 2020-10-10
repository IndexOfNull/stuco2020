package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indexofnull/stuco2020/models"
)

//GetAllStudents gathers every Student record and responds with a JSON representation of them.
func GetAllStudents(c *gin.Context) {
	var students []models.Student
	models.GetAllStudents(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudent(c *gin.Context) {
	var student models.Student
	id := uint32(IntParam(c, "id"))
	models.GetStudent(&student, id)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) { //Currently deletes all users for some reason?
	student := models.Student{}
	id := c.Params.ByName("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	rowsAffected, err := models.DeleteStudent(&student, uint32(idNumber))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else if rowsAffected > 0 {
		c.JSON(http.StatusOK, "successfully deleted user")
	} else {
		c.JSON(http.StatusNotModified, "user not found")
	}
}
