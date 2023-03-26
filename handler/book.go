package handler

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/src/book"
)

type BookHandler struct {
	pb.UnimplementedBookServiceServer
	service book.Service
}

func NewBookHandler(service book.Service) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (bh *BookHandler) ServiceInsertNewBook(ctx context.Context, payload *pb.CreateBookRequest) (*pb.BookStatusResponse, error) {
	err := bh.service.AddNewBook(ctx, payload)
	if err != nil {

		return nil, err
	}

	response := &pb.BookStatusResponse{
		Status:  "Success",
		Message: "Success to input new book",
	}

	return response, nil
}

func (bh *BookHandler) ServiceFindBookById(ctx context.Context, payload *pb.FindBookByIdRequest) (*pb.BookStatusResponse, error) {
	findBook, err := bh.service.FindBookByID(ctx, payload.Id)
	if err != nil {

		return nil, err

	}

	book := &pb.BookResponse{
		Id:        findBook.ID,
		Title:     findBook.Title,
		Author:    findBook.Author,
		CreatedAt: findBook.CreatedAt.String(),
		UpdatedAt: findBook.UpdatedAt.String(),
	}

	response := &pb.BookStatusResponse{
		Status:  "Success",
		Message: "Book Found",
		Data:    book,
	}

	return response, nil
}




