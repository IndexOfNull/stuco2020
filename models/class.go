package models

import "github.com/indexofnull/stuco2020/config"

//Class represents a class. It is used to deliver the proper ballot for each student.
type Class struct {
	ID       uint32 `json:"id" gorm:"primaryKey;"`
	Name     string `json:"string" gorm:"size:128;not null"`
	Global   bool   `json:"global" gorm:"default:0;not null"`
	NumVotes uint8  `json:"vote_count" gorm:"default:1;not null"`
	//Students []Student //`gorm:"foreignKey:ClassID"`
}

func (c *Class) TableName() string {
	return "classes"
}

func GetClassMembers(students *[]Student, id uint32) error {
	if err := config.DB.Model(&Student{}).Where("class_id = ?", id).Find(students).Error; err != nil {
		return err
	}
	return nil
}

func GetClass(class *Class, id uint32) error {
	if err := config.DB.Model(&Class{}).Where("id = ?", id).Find(class).Error; err != nil {
		return err
	}
	return nil
}
