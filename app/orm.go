package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db = GetOrm()
)

type BashHistory struct {
	gorm.Model
	Command string
	Path    string
	Output  string
	Status  int
}

type ProcessHistory struct {
	gorm.Model
	Pid      int
	Name     string
	Username string
	Action   string
}

func GetOrm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if (!db.Migrator().HasTable(&BashHistory{}) && !db.Migrator().HasTable(&ProcessHistory{})) {
		db.AutoMigrate(&BashHistory{}, &ProcessHistory{})
	}
	return db
}
