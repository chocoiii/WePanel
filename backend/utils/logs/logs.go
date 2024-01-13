package logs

import (
	"WePanel/backend/global"
	"WePanel/backend/utils/config"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var infoLogName = filepath.Join(global.BASEDIR, config.GetConfig("Log", "infoLog"))
var errorLogName = filepath.Join(global.BASEDIR, config.GetConfig("Log", "errorLog"))
var fatalLogName = filepath.Join(global.BASEDIR, config.GetConfig("Log", "fatalLog"))
var infoLink = filepath.Join(global.BASEDIR, config.GetConfig("Log", "infoLink"))
var errorLink = filepath.Join(global.BASEDIR, config.GetConfig("Log", "errorLink"))
var fatalLink = filepath.Join(global.BASEDIR, config.GetConfig("Log", "fatalLink"))

func newLfsHook() log.Hook {
	infoWriter, infoErr := rotatelogs.New(
		filepath.Join(infoLogName, "%Y_%m_%d.log"),
		rotatelogs.WithLinkName(infoLink),
		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),
		// WithMaxAge设置文件清理前的最长保存时间
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationSize(10*1024*1024),
	)
	if infoErr != nil {
		log.Errorf("Log config for info_history logger error: %v", infoErr)
	}
	errorWriter, errorErr := rotatelogs.New(
		filepath.Join(errorLogName, "%Y_%m_%d.log"),
		rotatelogs.WithLinkName(errorLink),
		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),
		// WithMaxAge和WithRotationCount二者只能设置一个
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationSize(10*1024*1024),
	)
	if errorErr != nil {
		log.Errorf("Log config error_history logger error: %v", errorErr)
	}
	fatalWriter, fatalErr := rotatelogs.New(
		filepath.Join(fatalLogName, "%Y_%m_%d.log"),
		rotatelogs.WithLinkName(fatalLink),
		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),
		// WithMaxAge和WithRotationCount二者只能设置一个
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationSize(10*1024*1024),
	)
	if fatalErr != nil {
		log.Errorf("Log config fatal_history logger error: %v", fatalErr)
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.InfoLevel:  infoWriter,
		log.ErrorLevel: errorWriter,
		log.FatalLevel: fatalWriter,
	}, &nested.Formatter{
		TimestampFormat: global.TimeFormat,
		NoColors:        true,
		HideKeys:        true,
		ShowFullLevel:   true,
	})
	return lfsHook
}

func Init() {
	ll := log.New()
	ll.SetFormatter(&nested.Formatter{
		TimestampFormat: global.TimeFormat,
		NoColors:        true,
		HideKeys:        true,
		ShowFullLevel:   true,
	})
	ll.SetOutput(os.Stdout)
	ll.AddHook(newLfsHook())
	ll.SetReportCaller(true)
	ll.SetLevel(log.InfoLevel)
	global.LOG = ll
	global.LOG.Info("init logger successfully!")
}
