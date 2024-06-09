package main

import (
	"pembiayaan/src/definitions"
	"pembiayaan/src/interfaces/http"

	"log"
	"os"

	gormDriver "pembiayaan/src/drivers/gorm"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	gorm, errGorm := gormDriver.Connect()
	appInterface := os.Getenv("INTERFACE")

	appContext := definitions.AppContext{
		Gorm: gorm,
	}

	if errGorm != nil {
		log.Fatalf("Failed to Initialized Database: %v", errGorm)
	}

	if appInterface == "" {
		log.Fatal("Interface not found")
	}

	if appInterface == "http" {
		http := http.NewHttp(&appContext)
		http.Launch()
	}

	if appInterface != "" {
		log.Fatalf(`Interface not found (%v)`, appInterface)
	}
}
