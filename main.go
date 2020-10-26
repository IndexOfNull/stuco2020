package main

import (
	"fmt"
	"log"

	"flag"

	"github.com/gin-gonic/gin"
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
var port = flag.String("p", "8080", "determines what port the program should bind to")

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
		config.DB.AutoMigrate(&models.Class{}, &models.Student{}, &models.Code{}, &models.Vote{})
	}

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

	gin.SetMode(gin.ReleaseMode)
	log.Println("Starting backend on :" + *port)

	r := routes.SetupRouter()
	err := r.Run(":" + *port)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
