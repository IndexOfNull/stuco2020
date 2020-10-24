package main

import (
	"fmt"

	"flag"

	"github.com/indexofnull/stuco2020/config"
	"github.com/indexofnull/stuco2020/models"
	"github.com/indexofnull/stuco2020/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var shouldSeed = flag.Bool("s", false, "if set, the database will be seeded with testing values")
var shouldSeedFromFile = flag.Bool("ff", false, "if -s is set and this is set, it will parse seed data from seeding.json")
var shouldMigrate = flag.Bool("m", false, "if set, the database will be auto-migrated")
var shouldDrop = flag.Bool("d", false, "if set, the database will be dropped before the program starts")

var err error

func main() {
	flag.Parse()
	config.DB, err = gorm.Open(mysql.Open(config.BuildDBConfig().DbURL()), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
		return
	}

	if *shouldDrop == true {
		models.DropAll()
	}

	if *shouldMigrate == true {
		config.DB.Debug().AutoMigrate(&models.Class{}, &models.Student{}, &models.Code{}, &models.Vote{})
	}
	//migration.ServiceAutoMigration(config.DB)
	if *shouldSeed == true {
		if *shouldSeedFromFile {
			err := models.SeedFromFile()
			if err != nil {
				panic(err)
			}
		} else {
			models.Seed()
		}
	}

	r := routes.SetupRouter()
	r.Run()
}
