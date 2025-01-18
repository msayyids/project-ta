package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitializeCloudinary() *cloudinary.Cloudinary {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")

	if apiKey == "" || apiSecret == "" || cloudName == "" {
		log.Panic("Cloudinary environment variables are not set properly")
	}

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Panicf("Failed to initialize Cloudinary: %v", err)
	}

	return cld
}
