package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	BASEDIR, _ = os.Getwd()
	CONFIGDIR  = filepath.Join(BASEDIR, "backend", "config", "config.ini")
	LOG        *logrus.Logger
	DB         *gorm.DB
)
