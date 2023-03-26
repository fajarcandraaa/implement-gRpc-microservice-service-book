package bookrepo

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/config/app"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity/bookentity"
	"github.com/pkg/errors"
)

func (b *BookRepository) FindBook(ctx context.Context, bookID string) (*bookentity.Book, error) {
	var book bookentity.Book
	err := b.db.First(&book, "id = ? ", bookID).Error
	if err != nil {
		parsed := b.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return nil, entity.ErrBookNotExist
		case app.ErrUniqueViolation:
			return nil, entity.ErrBooksCredentialNotExist
		default:
			return nil, errors.Wrap(parsed, "build statement query to find user from database")
		}
	}
	return &book, nil
}
