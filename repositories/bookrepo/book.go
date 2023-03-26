package bookrepo

import (
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/helpers/errorcodehandling"
	"github.com/jinzhu/gorm"
)

type BookRepository struct {
	db *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db:        db,
	}
}