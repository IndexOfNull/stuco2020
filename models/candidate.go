package models

type Candidate struct {
	ID        uint32
	StudentID uint32   `gorm:"not null"`
	Student   *Student `gorm:"constraint:OnDelete:CASCADE;"`
	/*ClassID   *uint    `gorm:"type:INTEGER;"`
	Class     *Class   `gorm:"constraint:OnDelete:SET NULL;"`*/
}

func (c *Candidate) TableName() string {
	return "candidates"
}
