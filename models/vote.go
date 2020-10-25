package models

import "time"

type Vote struct {
	ID        uint32    `json:"id" gorm:"primaryKey;"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	CodeID    *uint32   `json:"code_id" gorm:"not null"`
	Code      *Code     `gorm:"constraint:OnDelete:CASCADE"`
	StudentID *uint32   `json:"student_id" gorm:"not null;"`
	Student   *Student  `gorm:"constraint:OnDelete:CASCADE"`
}
