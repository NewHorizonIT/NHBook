package global

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config     config.Config
	MySQL      *gorm.DB
	Viper      *viper.Viper
	Logger     *zap.Logger
	Redis      *redis.Client
	Cloudinary *cloudinary.Cloudinary
)
