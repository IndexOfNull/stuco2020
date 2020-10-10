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
	Used      bool       `json:"used" gorm:"not null;default:0"`
}

//ResolvedCode represents a code that has been resolved with its corresponding information.
type ResolvedCode struct {
	Code *Code `json:"code"`
	//Student    *Student   `json:"student"`
	Candidates *[]Student `json:"candidates"`
}

func ResolveCode(resolved *ResolvedCode, code string) error {
	var co Code
	var candidates []Student
	if err := config.DB.Model(&co).Where("code = ?", code).Joins("Student").Preload("Student.Class").First(&co).Error; err != nil {
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

	*resolved = ResolvedCode{
		Code: &co,
		//Student:    co.Student,
		Candidates: &candidates,
	}

	fmt.Println(co)
	return nil
}
