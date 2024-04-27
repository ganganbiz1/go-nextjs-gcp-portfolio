package repository

import (
	"fmt"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/config"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/infra"
	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/logger"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type PrimaryDB struct {
	*gorm.DB
}

type ReplicaDB struct {
	*gorm.DB
}

func NewPrimaryDBConnect(c *config.PrimaryDBConfig) (PrimaryDB, func(), error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", c.Host, c.User, c.Password, c.Database, c.Port, c.SSLMode, c.TimeZone)

	db, err := gormtrace.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(c.LogLevel),
	}, gormtrace.WithServiceName("primaryDB"))

	// Datadog使わないときは以下を使用
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: gormLogger.Default.LogMode(c.LogLevel),
	// })

	if err != nil {
		return PrimaryDB{}, nil, infra.HandleError(err)
	}

	d, err := db.DB()

	if err != nil {
		return PrimaryDB{}, nil, infra.HandleError(err)
	}

	err = d.Ping()

	if err != nil {
		return PrimaryDB{}, nil, infra.HandleError(err)
	}
	cl := func() {
		d.Close()
	}
	logger.Info("PrimaryDB Connect Success")
	return PrimaryDB{db}, cl, nil
}

func NewReplicaDBConnect(c *config.ReplicaDBConfig) (ReplicaDB, func(), error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", c.Host, c.User, c.Password, c.Database, c.Port, c.SSLMode, c.TimeZone)

	db, err := gormtrace.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(c.LogLevel),
	}, gormtrace.WithServiceName("replicaDB"))

	// Datadog使わないときは以下を使用
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: gormLogger.Default.LogMode(c.LogLevel),
	// })

	if err != nil {
		return ReplicaDB{}, nil, infra.HandleError(err)
	}

	d, err := db.DB()

	if err != nil {
		return ReplicaDB{}, nil, infra.HandleError(err)
	}

	err = d.Ping()

	if err != nil {
		return ReplicaDB{}, nil, infra.HandleError(err)
	}
	cl := func() {
		d.Close()
	}
	logger.Info("ReplicaDB Connect Success")
	return ReplicaDB{db}, cl, nil
}
