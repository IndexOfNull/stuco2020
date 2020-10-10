package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//PongHandler sends a friendly "pong!" message back
func PongHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}

//Converts the specified parameter to an integer, returning a 400 (bad request) error if the input is invalid.
func IntParam(c *gin.Context, param string) int {
	n, err := strconv.Atoi(c.Params.ByName(param))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	return n
}
