package models

import (
	"github.com/indexofnull/stuco2020/config"
)

//Students
var students []Student

var people = []map[string]interface{}{
	//HS
	{"FirstName": "Damien", "LastName": "Kazewych", "ClassID": 1, "Candidate": 1},
	{"FirstName": "Daniel", "LastName": "Kazewych", "ClassID": 1},
	{"FirstName": "Kush", "LastName": "Patel", "ClassID": 1},
	{"FirstName": "Grant", "LastName": "Sears", "ClassID": 1, "Candidate": 1},
	//MS
	{"FirstName": "Harry", "LastName": "Something", "ClassID": 2},
	{"FirstName": "Emery", "LastName": "Swank", "ClassID": 2, "Candidate": 1},
	{"FirstName": "Johan", "LastName": "Delgado", "ClassID": 2, "Candidate": 1},
	{"FirstName": "Mary", "LastName": "Roberts", "ClassID": 2},
	//Teachers
	{"FirstName": "Tim", "LastName": "Garrett", "ClassID": nil},
	{"FirstName": "Casy", "LastName": "Smith", "ClassID": nil},
	{"FirstName": "Eileen", "LastName": "Fiarbanks", "ClassID": nil},
	{"FirstName": "David", "LastName": "Mathews", "ClassID": nil},
	//4th Grader
	{"FirstName": "Star", "LastName": "Knight", "ClassID": 3, "Candidate": 1},
}

var classes = []map[string]interface{}{
	{"ID": 1, "Name": "High School", "Global": true},
	{"ID": 2, "Name": "Middle School", "Global": false},
	{"ID": 3, "Name": "4th Grade", "Global": false},
}

var codes = []map[string]interface{}{
	{"StudentID": 1, "Code": "1234"},
	{"StudentID": 5, "Code": "abcd"},
	{"StudentID": 9, "Code": "asdf"},
}

//Seed populates the database with predefined values
func Seed() {
	config.DB.Model(&Class{}).Create(&classes)
	for i := 0; i < 1; i++ {
		config.DB.Model(&Student{}).Create(&people)
	}
	config.DB.Model(&Code{}).Create(&codes)
}

func DropAll() {
	config.DB.Migrator().DropTable("classes", "students", "codes", "votes")
}
