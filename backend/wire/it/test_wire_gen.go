// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package it

import (
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler/validate"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/repository"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/wire/it/di"
)

// Injectors from test_wire.go:

func DI() (*DIManager, func(), error) {
	errorHandler := handler.NewErrorHandler()
	customValidator := validate.NewCustomValidator()
	testServer := di.NewTestServer(errorHandler, customValidator)
	primaryDB, cleanup, err := di.NewTestPrimaryDBConnect()
	if err != nil {
		return nil, nil, err
	}
	baseRepository := repository.NewBaseRepository(primaryDB)
	replicaDB, cleanup2, err := di.NewTestReplicaDBConnect()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	ifUserRepository := repository.NewUserRepository(baseRepository, replicaDB)
	testContainer := di.NewTestContainer(testServer, baseRepository, primaryDB, replicaDB, ifUserRepository)
	diManager := &DIManager{
		DIC: testContainer,
	}
	return diManager, func() {
		cleanup2()
		cleanup()
	}, nil
}

// test_wire.go:

type DIManager struct {
	DIC *di.TestContainer
}