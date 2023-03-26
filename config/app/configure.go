package app

import (
	"github.com/fajarcandraaa/implement-gRpc-microservice-service-book/internal/entity/bookentity"
)

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&bookentity.Book{},
	}

	return migrationData
}
