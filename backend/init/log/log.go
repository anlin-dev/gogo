package log

import (
	"github.com/sirupsen/logrus"
	"gogo/backend/global"
)

func Init() {
	l := logrus.New()
	global.LOG = l
	global.LOG.Info("init logger successfully")
}
