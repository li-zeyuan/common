package mysqlstore

import (
	"time"

	"github.com/li-zeyuan/common/mylogger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

var Db *gorm.DB

func New(conf *Config) error {
	if Db != nil {
		return nil
	}

	var err error
	Db, err = gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: buildLogger(),
	})
	if err != nil {
		return err
	}
	sqlDb, err := Db.DB()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(conf.MaxConn)
	sqlDb.SetMaxOpenConns(conf.MaxOpen)
	sqlDb.SetConnMaxIdleTime(time.Duration(conf.Timeout))
	return nil
}

func buildLogger() zapgorm2.Logger {
	logger := zapgorm2.New(mylogger.GetZapLogger())
	logger.SlowThreshold = time.Second
	//logger.IgnoreRecordNotFoundError = true
	logger.LogLevel = gormlogger.Info
	//logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	return logger
}

func Close() {
	sqlDb, err := Db.DB()
	if err != nil {
		return
	}

	sqlDb.Close()
}
