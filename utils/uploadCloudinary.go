package utils

import (
	"context"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Cloudinary config
func StartCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinary_url := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(cloudinary_url)
	return cld, err
}

// uploads file to Cloudinary and returns the public URL
func UploadToCloudinary(file *multipart.FileHeader, folder string) (string, error) {
	// open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// upload options
	uploadOptions := uploader.UploadParams{
		Folder: folder,
	}

	// upload file
	cloudinaryURL, err := uploadFile(src, file.Filename, uploadOptions)
	if err != nil {
		return "", err
	}

	return cloudinaryURL, nil
}

func uploadFile(file multipart.File, filename string, uploadOptions uploader.UploadParams) (string, error) {
	// initialize cloudinary
	c, err := StartCloudinary()
	if err != nil {
		return "", err
	}

	// create a context with timeout
	ctx := context.Background()

	// upload file to cloudinary
	uploadResult, err := c.Upload.Upload(ctx, file, uploadOptions)
	if err != nil {
		return "", err
	}

	// return the public URL
	return uploadResult.SecureURL, nil
}