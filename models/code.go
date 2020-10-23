package models

import (
	"time"

	"github.com/indexofnull/stuco2020/config"
)

//Code represents a voting code needed to initialize a vote.
type Code struct {
	ID        uint32     `json:"id" gorm:"primaryKey;"`
	StudentID uint32     `json:"student_id"`
	Student   *Student   `json:"student" gorm:"constraint:OnDelete:SET NULL"`
	Code      string     `json:"code" gorm:"size:16;not null;index;"`
	CreatedAt *time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP;"`
	UpdatedAt *time.Time `json:"updated_at"`
	TimesUsed uint16     `json:"times_used" gorm:"not null;default:0"`
	MaxUses   uint16     `json:"max_uses" gorm:"not null;default:1"`
	Active    bool       `json:"active" gorm:"not null;default:1"`
}

/*func GetCodeByCode(c *Code, code string) error {
	if err := config.DB.Where("code = ?", code).Joins("Student").Preload("Student.Class").First(&c).Error; err != nil {
		return err
	}
	return nil
}*/

/*func GetCodeByID(c *Code, code uint32) error {
	if err := config.DB.Where("id = ?", code).Joins("Student").Preload("Student.Class").First(&c).Error; err != nil {
		return err
	}
	return nil
}*/

func ResolveCode(resolved *Code, code string) error {
	var co Code
	//var candidates []Student
	if err := config.DB.Preload("Student.VotesFor.Students", "candidate = ?", true).Where("code = ?", code).Joins("Student").Find(&co).Error; err != nil {
		return err
	}

	*resolved = co
	return nil
}
