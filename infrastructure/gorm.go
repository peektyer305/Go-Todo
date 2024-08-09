package infrastructure

import (
	"fmt"
	"sync"
	"time"

	"github.com/peektyer305/Go-Todo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var gormPostgresInstance *gorm.DB
var gormPostgresOnce sync.Once

func NewGormPostgres() *gorm.DB {
	gormPostgresOnce.Do(func() {
		logMode := logger.Silent
		if config.Conf.IsLocal() {
			logMode = logger.Info
		}
		db, err := gorm.Open(
			postgres.New(postgres.Config{
				DSN:                  getDbDsn(),
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			}),
			&gorm.Config{
				Logger: logger.Default.LogMode(logMode),
			},
		)
		if err != nil {
			panic(err)
		}
		conn, err := db.DB()
		if err != nil {
			panic(err)
		}
		conn.SetConnMaxIdleTime(time.Hour)
		conn.SetConnMaxLifetime(24 * time.Hour)
		conn.SetMaxIdleConns(100)
		conn.SetMaxOpenConns(200)

		gormPostgresInstance = db.Preload(clause.Associations).Session(&gorm.Session{FullSaveAssociations: true})
	})
	return gormPostgresInstance
}

func getDbDsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", config.Conf.Db.Host, config.Conf.Db.User, config.Conf.Db.Password, config.Conf.Db.Name, config.Conf.Db.Port)
}
