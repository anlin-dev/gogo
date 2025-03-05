package global

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	LOG *logrus.Logger

	Viper *viper.Viper
)
