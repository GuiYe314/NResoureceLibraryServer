package db

import (
	"NResoureceLibraryServer/db/gormLink"

	"gorm.io/gorm"
)

func Init_DB(dbConfig *gormLink.DBConfig) (*gorm.DB, error) {
	return gormLink.LinkDBType(dbConfig)
}
