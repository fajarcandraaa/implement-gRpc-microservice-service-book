package bookentity

import "time"

type Book struct {
	ID        string    `gorm:"size:36;not null;unique index;primaryKey"`
	Title     string    `gorm:"size:255;unique"`
	Author    string    `gorm:"size:100;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
