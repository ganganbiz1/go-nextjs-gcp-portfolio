package di

import (
	"fmt"
	"os"

	domain_repo "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain/repository"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/handler/validate"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra/repository"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type TestContainer struct {
	TestServer     *TestServer
	BaseRepository repository.BaseRepository
	PrimaryDB      repository.PrimaryDB
	ReplicaDB      repository.ReplicaDB
	UserRepository domain_repo.IfUserRepository
}

func NewTestContainer(
	testServer *TestServer,
	baseRepository repository.BaseRepository,
	primaryDB repository.PrimaryDB,
	replicaDB repository.ReplicaDB,
	userRepository domain_repo.IfUserRepository,
) *TestContainer {
	return &TestContainer{
		testServer,
		baseRepository,
		primaryDB,
		replicaDB,
		userRepository,
	}
}

type TestServer struct {
	Echo *echo.Echo
}

func NewTestServer(
	errHandler *handler.ErrorHandler,
	validator *validate.CustomValidator,
) *TestServer {
	e := echo.New()

	e.HTTPErrorHandler = errHandler.Handler
	e.Validator = validator

	return &TestServer{
		Echo: e,
	}
}

func NewTestPrimaryDBConnect() (repository.PrimaryDB, func(), error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST_TEST"),
		os.Getenv("POSTGRES_USER_TEST"),
		os.Getenv("POSTGRES_PASSWORD_TEST"),
		os.Getenv("POSTGRES_DB_TEST"),
		5432,
		os.Getenv("POSTGRES_SSL_MODE_TEST"),
		os.Getenv("POSTGRES_TIME_ZONE_TEST"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		return repository.PrimaryDB{}, nil, infra.HandleError(err)
	}

	d, err := db.DB()

	if err != nil {
		return repository.PrimaryDB{}, nil, infra.HandleError(err)
	}

	err = d.Ping()

	if err != nil {
		return repository.PrimaryDB{}, nil, infra.HandleError(err)
	}
	cl := func() {
		d.Close()
	}
	logger.Info("TestPrimaryDB Connect Success")
	return repository.PrimaryDB{DB: db}, cl, nil
}

func NewTestReplicaDBConnect() (repository.ReplicaDB, func(), error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST_TEST"),
		os.Getenv("POSTGRES_USER_TEST"),
		os.Getenv("POSTGRES_PASSWORD_TEST"),
		os.Getenv("POSTGRES_DB_TEST"),
		5432,
		os.Getenv("POSTGRES_SSL_MODE_TEST"),
		os.Getenv("POSTGRES_TIME_ZONE_TEST"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		return repository.ReplicaDB{}, nil, infra.HandleError(err)
	}

	d, err := db.DB()

	if err != nil {
		return repository.ReplicaDB{}, nil, infra.HandleError(err)
	}

	err = d.Ping()

	if err != nil {
		return repository.ReplicaDB{}, nil, infra.HandleError(err)
	}
	cl := func() {
		d.Close()
	}
	logger.Info("ReplicaDB Connect Success")
	return repository.ReplicaDB{DB: db}, cl, nil
}
