package routers

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/config/app"
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/handler"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

var grpcServer = grpc.NewServer()

type Service struct {
	Book *handler.BookHandler
}

type Serve struct {
	DB      *gorm.DB
	Router  *mux.Router
	Service Service
}

// Initialize is used to initialize db driver connection
func (s *Serve) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	//Set migration table
	registry := app.SetMigrationTable()

	//initiate database driver
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		s.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		s.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	//Migration proccess
	s.DB.Debug().AutoMigrate(registry...) //database migration

	s.Router = mux.NewRouter()

	s.initializeRoutes()
}

// Run is used to execute connection and run our service
//
//	func (s *Serve) Run() {
//		port := os.Getenv("APP_PORT")
//		fmt.Println("Listening to port ", port)
//		log.Fatal(http.ListenAndServe(":"+port, s.Router))
//	}
func (s *Serve) RunGrpc() {
	port := os.Getenv("GRPC_PORT")
	fmt.Println("Listening to port ", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed connection : %v\n", err)
	}

	log.Printf("server listening at %v\n", lis.Addr())
	if grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server : %v\n", err)
	}
	// log.Fatal(http.ListenAndServe(":"+port, s.Router))
}
