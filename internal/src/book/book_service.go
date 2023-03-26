package book

import (
	"context"
	"time"

	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/helpers/errorcodehandling"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/helpers/unique"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity/bookentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/repositories"
	"github.com/google/uuid"
)

type service struct {
	repo *repositories.Repository
	err  *errorcodehandling.CodeError
}

func NewService(repo *repositories.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) AddNewBook(ctx context.Context, payload *pb.CreateBookRequest) error {
	book := &bookentity.Book{
		ID:        uuid.NewString(),
		Title:     payload.Title,
		Author:    payload.Author,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err := s.repo.Book.InsertNewBook(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindBookByID(ctx context.Context, bookID string) (*bookentity.Book, error) {
	if err := unique.ValidateUUID(bookID); err != nil {
		return nil, entity.ErrBookNotExist
	}

	book, err := s.repo.Book.FindBook(ctx, bookID)
	if err != nil {
		return nil, err
	}

	return book, nil
}
