package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Source provides a facebook page as source
type Source struct {
	ID        uint
	UserName  string `gorm:"not null; unique"`
	CreatedAt time.Time
	Results   []Result `gorm:"ForeignKey:SourceID"`
}

// InsertSources inserts base source data
func InsertSources(db *gorm.DB) []*Source {
	sources := []*Source{
		&Source{
			ID:       1,
			UserName: "TheIDEAlistRevolution",
		},
	}

	for _, s := range sources {
		db.FirstOrCreate(s, map[string]interface{}{"id": s.ID})
	}

	return sources
}
