package utils

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImg(file multipart.File, fileName string, cld *cloudinary.Cloudinary) (string, error) {
	ctx := context.Background()

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: fileName,
		Folder:   "upload",
	})
	fmt.Println("URL: " + result.SecureURL)
	if err != nil {
		return "", err
	}

	return result.SecureURL, nil
}
