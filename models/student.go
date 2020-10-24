package models

import "github.com/indexofnull/stuco2020/config"

//Student represents a student
type Student struct {
	ID        uint32   `json:"id" gorm:"primaryKey;"`
	FirstName *string  `json:"firstname" gorm:"not null;type:TEXT;"`
	LastName  *string  `json:"lastname" gorm:"not null;type:TEXT;"`
	Email     *string  `json:"email" gorm:"type:TEXT;"`
	Class     *Class   `json:"class" gorm:"constraint:OnDelete:SET NULL"`
	VotesFor  *[]Class `json:"votes_for" gorm:"many2many:votes_for;constraint:OnDelete:CASCADE;"`
	ClassID   *uint32  `json:"class_id"`
	Candidate bool     `json:"candidate" gorm:"not null;default:0;index"`
}

//Classes   []Class `gorm:"many2many:StudentID:classes" json:"classes" `

func (s *Student) TableName() string {
	return "students"
}

//GetStudentsByRange selects a slice of students (like a page.)
func GetStudentsByRange(students *[]Student, limit int, offset int) error {
	if err := config.DB.Joins("Class").Limit(limit).Offset(offset).Find(students).Error; err != nil {
		return err
	}
	return nil
}

//GetAllStudents retrieves all students from the database. USE THIS WISELY!
func GetAllStudents(students *[]Student) error {
	if err := config.DB.Joins("Class").Find(students).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentsByIDs(students *[]Student, ids *[]uint32) error {
	if err := config.DB.Joins("Class").Find(students, ids).Error; err != nil {
		return err
	}
	return nil
}

func GetCandidatesByIDs(students *[]Student, ids *[]uint32) error {
	if err := config.DB.Joins("Class").Where("candidate = ?", true).Find(students, ids).Error; err != nil {
		return err
	}
	return nil
}

//GetStudent retrieves a student by ID
func GetStudent(student *Student, id uint32) error {
	if err := config.DB.Debug().Joins("Class").Where(&Student{ID: id}).Find(student).Error; err != nil {
		return err
	}
	return nil
}

//CreateStudent takes the passed Student struct and inserts it into the database.
func CreateStudent(student *Student) error {
	if err := config.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

//DeleteStudent takes the passed Student and deletes it (if found).
func DeleteStudent(student *Student, id uint32) (int64, error) {
	q := config.DB.Debug().Where("id = ?", id).Delete(student)
	if q.Error != nil {
		return 0, q.Error
	}
	return q.RowsAffected, nil
}

//UpdateStudent takes the passed Student and writes any differences to the databse.
func UpdateStudent(student *Student) error {
	if err := config.DB.Save(student).Error; err != nil {
		return err
	}
	return nil
}

/*func GetClassmates(students *[]Student, id uint) error {
	if err := config.DB.Where("").
}*/
