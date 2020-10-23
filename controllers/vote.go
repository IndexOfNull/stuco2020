package controllers

type VoteBody struct {
	IDs []uint32 `json:"ids" binding:"required"`
}

/*func CastVote(c *gin.Context) {
	var b VoteBody
	err := c.BindJSON(&b)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var code models.Code
	err = models.GetCodeByCode(&code, c.Params.ByName("code"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error while getting code"})
		return
	}

	if code.ID == 0 { //If the code is zero, which should not be possible with auto_increment under normal circumstances
		c.JSON(http.StatusNotFound, gin.H{"error": "code does not exist."})
		return
	} else if code.TimesUsed >= code.MaxUses {
		c.JSON(http.StatusForbidden, gin.H{"error": "code has been used too many times"})
		return
	} else if code.Active == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "code is not activated"})
		return
	} else if code.Student == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "code is orphaned and is therefore not valid; contact the code manager to give it a new owner."})
		return
	}

	var candidates []models.Student
	models.GetCandidatesByIDs(&candidates, &b.IDs) //Possible DOS attack if they give us a ton of IDs?
	fmt.Print(candidates)

	//.JSON(http.StatusOK, resolved)
}*/
