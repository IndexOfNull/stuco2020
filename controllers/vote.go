package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indexofnull/stuco2020/models"
)

func CastVote(c *gin.Context) {
	var resolved models.ResolvedCode
	models.ResolveCode(&resolved, c.Params.ByName("code"))
	c.JSON(http.StatusOK, resolved)
}
