package routers

import (
	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/handler"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/src/book"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/repositories"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	r := repositories.NewRepository(se.DB)

	//======================== ROUTER ========================
	//Setting Services
	//Setting User Service

	//=== USER ===
	bookService := book.NewService(r)
	bookHandler := handler.NewBookHandler(bookService)
	pb.RegisterBookServiceServer(grpcServer, bookHandler)
	//=========================================================

}
