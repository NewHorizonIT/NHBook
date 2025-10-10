package global

import (
	"database/sql"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Config     config.Config
	MySQL      *sql.DB
	Viper      *viper.Viper
	Logger     *zap.Logger
	Redis      *redis.Client
	Cloudinary *cloudinary.Cloudinary
)
