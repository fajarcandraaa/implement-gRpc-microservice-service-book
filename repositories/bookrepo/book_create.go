package bookrepo

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/config/app"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity/bookentity"
	"github.com/pkg/errors"
)

func (b *BookRepository) InsertNewBook(ctx context.Context, payload *bookentity.Book) error {
	err := b.db.Create(payload).Error
	if err != nil {
		parsed := b.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return entity.ErrBookNotExist
		case app.ErrUniqueViolation:
			return entity.ErrBookAlreadyExist
		default:
			return errors.Wrap(parsed, "build statement query to insert book from database")
		}
	}
	return nil
}
