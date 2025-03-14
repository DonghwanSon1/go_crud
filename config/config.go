package config

import (
	"github.com/joho/godotenv"
	"github.com/naoina/toml"
	"go_crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
	Jwt struct {
		Secret string
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}

var DB *gorm.DB

func ConnectToDb() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/golang"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect to the database")
	}
}

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
