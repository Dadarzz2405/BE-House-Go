package config

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
)

var CLD *cloudinary.Cloudinary

func ConnectCloudinary() {
	loadEnv()

	cloudName := getEnv("CLOUDINARY_CLOUD_NAME", "CLOUDINARY_NAME")
	apiKey := getEnv("CLOUDINARY_API_KEY", "CLOUDINARY_KEY")
	apiSecret := getEnv("CLOUDINARY_API_SECRET", "CLOUDINARY_SECRET")

	if cloudName == "" || apiKey == "" || apiSecret == "" {
		CLD = nil
		log.Println("cloudinary is not configured: set CLOUDINARY_CLOUD_NAME/CLOUDINARY_API_KEY/CLOUDINARY_API_SECRET or CLOUDINARY_NAME/CLOUDINARY_KEY/CLOUDINARY_SECRET")
		return
	}

	cld, err := cloudinary.NewFromParams(
		cloudName,
		apiKey,
		apiSecret,
	)
	if err != nil {
		CLD = nil
		log.Printf("failed to connect to cloudinary: %v", err)
		return
	}
	CLD = cld
}
