package book

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity/bookentity"
)

type Service interface {
	AddNewBook(ctx context.Context, payload *pb.CreateBookRequest) error
	FindBookByID(ctx context.Context, bookID string) (*bookentity.Book, error)
}
