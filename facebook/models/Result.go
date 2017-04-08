package models

type Result struct {
	ID            uint
	PictureRawURL string `gorm:"unique"`
	FbID          string `gorm:"unique"`
	Description   string
	SourceID      uint
}
