package db

import (
	"invokes/internal/config"
	"invokes/internal/utils"
	"time"

	gorm_logrus "github.com/onrik/gorm-logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Wrapper in case of had to choose between multiple orm
type Wrapper struct {
	GormDB *gorm.DB
}

// Initialize create the associated db
func (db *Wrapper) Initialize(config *config.Config) error {

	var err error
	switch config.DatabaseConfig.Engine {
	case "mysql":
		utils.Logger.Debugf("Creating mysqldatabase at %s", config.DatabaseConfig.ConnectionString)
		db.GormDB, err = gorm.Open(mysql.Open(config.DatabaseConfig.ConnectionString), &gorm.Config{
			Logger:                 gorm_logrus.New(),
			SkipDefaultTransaction: true,
		})
	case "sqlite":
		utils.Logger.Debugf("Creating sqlitedatabase at %s", config.DatabaseConfig.ConnectionString)
		db.GormDB, err = gorm.Open(sqlite.Open(config.DatabaseConfig.ConnectionString), &gorm.Config{
			Logger:                 gorm_logrus.New(),
			SkipDefaultTransaction: true,
		})
	default:
		utils.Logger.Errorf("No database %s", config.DatabaseConfig.Engine)
	}

	if db.GormDB == nil {
		utils.Logger.Errorf("1 Can't connect to %s due to %s", config.DatabaseConfig.ConnectionString, err)
		return err
	}

	if err != nil {
		utils.Logger.Errorf("2 Can't connect to %s due to %s", config.DatabaseConfig.ConnectionString, err)
		return err
	}

	sqlDB, err := db.GormDB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetConnMaxLifetime(config.DatabaseConfig.ConnMaxLifetime * time.Second)
	sqlDB.SetMaxIdleConns(config.DatabaseConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.DatabaseConfig.MaxOpenConns)

	retries := 0
	for retries < 30 {
		err = sqlDB.Ping()
		if err == nil {
			break
		}
		utils.Logger.Infof("Failed to ping database : %s (retry %d)", config.DatabaseConfig.ConnectionString, retries)
		time.Sleep(time.Second)
		retries++
	}
	return err
}
