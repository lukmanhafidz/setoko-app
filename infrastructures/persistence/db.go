package persistence

import (
	"fmt"
	"log"
	"os"
	"setokoapp/domain/model"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnectDb() (*gorm.DB, error) {
	config := model.Config.Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta",
		config.Host, config.User, config.Password, config.Name, config.Port) //get value from config file

	gormCfg := new(gorm.Config)

	gormLog := logger.New( //log sql queries
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //set up threshold for slow sql queries
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	gormCfg.Logger = gormLog

	db, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		return nil, err
	}

	db.NamingStrategy = schema.NamingStrategy{SingularTable: true}
	return db.Session(&gorm.Session{PrepareStmt: true}), nil
}
