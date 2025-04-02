package initializer

import (
	"log"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/global"
	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinary() {
	cldCnf := global.Config.Cloudinary
	cld, err := cloudinary.NewFromParams(cldCnf.CloudName, cldCnf.ApiKey, cldCnf.ApiSecret)

	if err != nil {
		log.Fatalf("Connect Cloudinary Error: %v", err)
	}
	global.Cloudinary = cld
}
