package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indexofnull/stuco2020/config"
	"github.com/indexofnull/stuco2020/models"
)

type VoteBodyClass struct {
	ID       uint32   `json:"id" binding:"required"`
	Selected []uint32 `json:"selected" binding:"required"`
}
type VoteBody struct {
	Ballot []VoteBodyClass `json:"ballot" binding:"required"`
}

func CastVote(c *gin.Context) {
	//Get the request body
	var b VoteBody
	err := c.BindJSON(&b)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//Get the code from the URL, error if not found
	codeParam := c.Params.ByName("code")
	if codeParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "code parameter not specified"})
		return
	}

	//Resolve the code
	var code models.Code
	err = models.ResolveCode(&code, codeParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	//Checks about the code
	if code.ID == 0 { //Assume that 0 ids aren't going to be a thing
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "code does not exist"})
		return
	}
	if code.Active == false {
		c.JSON(http.StatusForbidden, gin.H{"status": "error", "error": "code is not active"})
		return
	}
	if code.TimesUsed >= code.MaxUses {
		c.JSON(http.StatusForbidden, gin.H{"status": "error", "error": "code has no more uses left"})
		return
	}

	var finalVotes []*models.Vote
	for _, class := range b.Ballot { //Checks about the ballot

		//Deep integrity checks
		if class.Selected == nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": fmt.Sprintf("selected field not included for class %d", class.ID)})
			return
		}

		//Resolve the given class ID to a class, error if not found
		var referenceClass *models.Class
		for _, rC := range *code.Student.VotesFor { //Pretty sure this is a big no-no but whatever
			if rC.ID == uint32(class.ID) {
				referenceClass = &rC
				break
			}
		}
		if referenceClass == nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": fmt.Sprintf("class %d doesn't exist or is outside of code scope", class.ID)})
			return
		}

		//Check for non-empty selections that aren't completely filled out (abstain if 0)
		if len(class.Selected) != 0 && len(class.Selected) != int(referenceClass.NumVotes) {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": fmt.Sprintf("invalid selection count for class %d (must have 0 or %d)", class.ID, referenceClass.NumVotes)})
			return
		}

		//Check for invalid candidates in selection
		uniqueStudentIDs := make([]uint32, len(class.Selected))
		for idx, studentID := range class.Selected {

			//Check to see if we've seen this student before
			for _, uniqueID := range uniqueStudentIDs {
				if uniqueID == studentID {
					c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": fmt.Sprintf("student %d appears more than once in class %d", studentID, class.ID)})
					return
				}
			}
			uniqueStudentIDs[idx] = studentID //We've seen it now, so toss it in the lsit

			//Check if they're even in the class
			studentInClass := false
			for _, referenceStudentID := range referenceClass.Students {
				if studentID == referenceStudentID.ID {
					studentInClass = true
				}
			}
			if studentInClass == false {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": fmt.Sprintf("student %d not in class %d", studentID, referenceClass.ID)})
				return
			}

			//Append a vote to the slice
			v := &models.Vote{
				CodeID:    &code.ID,
				StudentID: &studentID,
			}
			finalVotes = append(finalVotes, v)
		}

	}
	if len(finalVotes) > 0 {
		code.TimesUsed++ //Technically you could race condition votes in order to get more than one vote in

		err := config.DB.Save(&code).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "could not increment code use counter, discarding vote"})
			return
		}

		result := config.DB.Create(&finalVotes)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": result.Error.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
