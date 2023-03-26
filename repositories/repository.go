package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Book Book
}

//NewRepository to setting services repositories
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Book: NewBook(db),
	}
}
