
package graph

import (
	"database/sql"
	"log"
	"time"

	"backend/graph/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "modernc.org/sqlite" 
)


var DB *gorm.DB

func InitDB() *gorm.DB {
	sqlDB, err := sql.Open("sqlite", "expenses.db")
	if err != nil {
		log.Fatalf("failed to open SQLite DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB, err = gorm.Open(sqlite.Dialector{Conn: sqlDB}, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to initialise Gorm: %v", err)
	}

	if err := DB.AutoMigrate(&model.Expense{}, &model.InternalUser{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	return DB
}
