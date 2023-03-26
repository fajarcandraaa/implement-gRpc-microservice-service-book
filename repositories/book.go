package repositories

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity/bookentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/repositories/bookrepo"
	"github.com/jinzhu/gorm"
)

type Book interface {
	InsertNewBook(ctx context.Context, payload *bookentity.Book) error
	FindBook(ctx context.Context, bookID string) (*bookentity.Book, error)
}

func NewBook(db *gorm.DB) Book {
	return bookrepo.NewBookRepository(db)
}
