package models

import (
	"github.com/indexofnull/stuco2020/config"
)

//Students
var students []Student

var s1 = "Damien"
var s2 = "Kazewych"
var i1 uint32 = 1

var people = []Student{
	{
		FirstName: &s1,
		LastName:  &s2,
		VotesFor:  &[]Class{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}},
		ClassID:   &i1,
		Candidate: true,
	},
}

var classes = []Class{
	{
		Name: "High School",
	},
	{
		Name: "Middle School",
	},
	{
		Name: "10th Grade",
	},
	{
		Name: "Teachers",
	},
}

var codes = []Code{
	{
		Code:      "damien",
		StudentID: 1,
	},
}

//Seed populates the database with predefined values
func Seed() {
	config.DB.Model(&Class{}).Create(&classes)
	config.DB.Debug().Model(&Student{}).Create(&people)
	config.DB.Model(&Code{}).Create(&codes)
}
	}
	config.DB.Model(&Code{}).Create(&codes)
}

func DropAll() {
	config.DB.Migrator().DropTable("classes", "students", "codes", "votes")
}
