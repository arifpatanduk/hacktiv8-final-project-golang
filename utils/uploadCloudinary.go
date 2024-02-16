package utils

import (
	"context"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Cloudinary config
func startCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinary_url := os.Getenv("CLOUDINARY_URL")
	cld, err := cloudinary.NewFromURL(cloudinary_url)
	return cld, err
}

// UPLOAD FILE
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
	c, err := startCloudinary()
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


// DELETE FILE
func DeleteFromCloudinary(publicURL string) error {
	// initialize Cloudinary
	c, err := startCloudinary()
	if err != nil {
		return err
	}

	// extract public URL to extract public ID
	publicID, err := getPublicIDFromURL(publicURL)
	if err != nil {
		return err
	}

	// Create a context with timeout
	ctx := context.Background()

	// Delete file from Cloudinary
	_, err = c.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID})
	return err
}

func getPublicIDFromURL(publicURL string) (string, error) {
	parsedURL, err := url.Parse(publicURL)
	if err != nil {
		return "", err
	}

	// extract the path and splitting it by "/"
	pathParts := strings.Split(parsedURL.Path, "/")

	// get the pathParts after base path
	publicIDSlice := pathParts[5:] // base path is usually until index 4
	filename := strings.TrimSuffix(publicIDSlice[len(publicIDSlice)-1], filepath.Ext(publicIDSlice[len(publicIDSlice)-1]))

	var publicID string
	if len(publicIDSlice) > 1 {
		for i := 0; i < len(publicIDSlice) - 1; i++ {
			publicID += publicIDSlice[i] + "/"
		}
	}
	publicID += filename

	return publicID, nil
}
