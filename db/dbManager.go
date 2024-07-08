package db

import (
	"NResoureceLibraryServer/db/gorm"
)

func Init_DB(dbConfig *gorm.DBConfig) {
	gorm.LinkDBType(dbConfig)
}
