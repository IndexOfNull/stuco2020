package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

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

type Seeds struct {
	Students []Student `json:"students"`
	Classes  []Class   `json:"classes"`
	Codes    []Code    `json:"codes"`
}

func SeedFromFile() error {
	ex, err := os.Executable() // the current executable
	if err != nil {
		return err
	}
	exPath := filepath.Dir(ex)
	data, err := ioutil.ReadFile(filepath.Join(exPath, "seed.json"))
	if err != nil {
		return err
	}

	var seeds Seeds
	if err = json.Unmarshal(data, &seeds); err != nil {
		return err
	}

	config.DB.Model(&Class{}).Create(&seeds.Classes)
	config.DB.Debug().Model(&Student{}).Create(&seeds.Students)
	config.DB.Model(&Code{}).Create(&seeds.Codes)
	return nil
}

func DropAll() {
	config.DB.Migrator().DropTable("classes", "students", "codes", "votes", "votes_for")
}
