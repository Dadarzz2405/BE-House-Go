package config

import (
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

var CLD *cloudinary.Cloudinary

func ConnectCloudinary() {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		panic("failed to connect to cloudinary")
	}
	CLD = cld
}
