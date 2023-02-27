package mysqlstore

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func New(conf *Config)  error {
	if Db != nil {
		return nil
	}

	var err error
	Db, err = gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
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

func Close() {
	sqlDb, err := Db.DB()
	if err != nil {
		return
	}

	sqlDb.Close()
}
