package handler_test

import (
	"fmt"

	test_wire "github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/wire/it"
	"gorm.io/gorm"
)

func di() (*test_wire.DIManager, func(), error) {
	di, cleanup, err := test_wire.DI()
	if err != nil {
		return nil, nil, err
	}
	return di, cleanup, err
}

func getDB(di *test_wire.DIManager) *gorm.DB {
	return di.DIC.PrimaryDB.DB
}

func truncate(db *gorm.DB) error {
	ts := []string{
		"articles",
		"users",
	}

	for _, t := range ts {
		if err := db.Exec(fmt.Sprintf("truncate table %s CASCADE", t)).Error; err != nil {
			return err
		}
	}

	return nil
}
