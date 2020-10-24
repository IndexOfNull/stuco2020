package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "code parameter not specified"})
		return
	}

	//Resolve the code
	var code models.Code
	err = models.ResolveCode(&code, codeParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, class := range b.Ballot { //Go through the ballot

		//Deep integrity checks
		if class.Selected == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("selected field not included for class %d", class.ID)})
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
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("class %d doesn't exist or is outside of code scope", class.ID)})
			return
		}
		fmt.Println(referenceClass.ID)

		//Check for non-empty selections that aren't completely filled out (abstain if 0)
		if len(class.Selected) != 0 && len(class.Selected) != int(referenceClass.NumVotes) {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("not enough selections made for class %d (must have 0 or %d)", class.ID, referenceClass.NumVotes)})
			return
		}

		//Check for invalid candidates in selection
		uniqueStudentIDs := make([]uint32, len(class.Selected))
		for idx, studentID := range class.Selected {

			//Check to see if we've seen this student before
			for _, uniqueID := range uniqueStudentIDs {
				if uniqueID == studentID {
					c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("student %d appears more than once in class %d", studentID, class.ID)})
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
				c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("student %d not in class %d", studentID, referenceClass.ID)})
				return
			}
		}

	}
	c.JSON(http.StatusOK, b)

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
