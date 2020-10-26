package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"flag"

	"github.com/gin-gonic/gin"
	"github.com/indexofnull/stuco2020/config"
	"github.com/indexofnull/stuco2020/models"
	"github.com/indexofnull/stuco2020/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var shouldSeed = flag.Bool("s", false, "if set, the database will be seeded with testing values")
var shouldSeedFromFile = flag.Bool("ff", false, "if -s is set and this is set, it will parse seed data from seeding.json")
var shouldMigrate = flag.Bool("m", false, "if set, the database will be auto-migrated")
var shouldDrop = flag.Bool("d", false, "if set, the database will be dropped before the program starts")

var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file, will try to use variables instead")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPortString := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbPort, err := strconv.Atoi(dbPortString)
	if err != nil {
		log.Fatalln("Could not parse database port. Please ensure that it is a valid number")
	}

	serverPort := os.Getenv("SERVER_PORT")
	fmt.Print(dbPass)
	if dbHost == "" || dbPort == 0 || dbUser == "" || dbName == "" {
		log.Fatalln("Missing databse information in .env file or environment variables")
	}

	if serverPort == "" {
		serverPort = "8080"
	}

	flag.Parse()
	config.DB, err = gorm.Open(mysql.Open(config.BuildDBConfig(dbHost, dbPort, dbUser, dbPass, dbName).DbURL()), &gorm.Config{})

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
	log.Println("Starting backend on :" + serverPort)

	r := routes.SetupRouter()
	err = r.Run(":" + serverPort)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
