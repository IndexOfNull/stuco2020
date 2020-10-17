package models

import (
	"fmt"
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

//ResolvedCode represents a code that has been resolved with its corresponding information.
type ResolvedCode struct {
	Code *Code `json:"code"`
	//Student    *Student   `json:"student"`
	Classes    map[uint32]Class `json:"classes"`
	Candidates *[]Student       `json:"candidates"`
}

func GetCodeByCode(c *Code, code string) error {
	if err := config.DB.Where("code = ?", code).Joins("Student").Preload("Student.Class").First(&c).Error; err != nil {
		return err
	}
	return nil
}

func GetCodeByID(c *Code, code uint32) error {
	if err := config.DB.Where("id = ?", code).Joins("Student").Preload("Student.Class").First(&c).Error; err != nil {
		return err
	}
	return nil
}

func ResolveCode(resolved *ResolvedCode, code string) error {
	var co Code
	var candidates []Student
	err := GetCodeByCode(&co, code)

	if err != nil {
		return err
	}

	if co.Student.Class == nil { //If the student doesn't belong to a class. Have them vote for everyone.
		if err := config.DB.Debug().Model(&candidates).Joins("Class").Where("candidate = ?", 1).Find(&candidates).Error; err != nil {
			return err
		}
	} else { //If they belong to a class, find candidates that aren't in classes, in the same class, or are in a global class
		if err := config.DB.Debug().Model(&candidates).Joins("Class").Where(
			config.DB.Where("students.class_id = ?", co.Student.ClassID).Or("students.class_id IS NULL").Or("Class.global = ?", true),
		).Where("students.candidate = ?", 1).Find(&candidates).Error; err != nil {
			return err
		}
	}

	//Build a class lookup table so that the client doesn't have to assemble one and we can save on payload size
	classes := make(map[uint32]Class)
	for i, cand := range candidates {
		if _, ok := classes[cand.Class.ID]; !ok {
			classes[cand.Class.ID] = *cand.Class //Dereference so we don't end up with a nil pointer when we clear it from the student
		}
		candidates[i].Class = nil
	}

	*resolved = ResolvedCode{
		Code:       &co,
		Classes:    classes,
		Candidates: &candidates,
	}

	fmt.Println(co)
	return nil
}
