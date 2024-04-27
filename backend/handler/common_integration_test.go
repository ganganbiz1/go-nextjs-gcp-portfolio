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
		"users",
	}
	// 外部キー制約のエラー出ないように設定
	if err := db.Exec("SET session_replication_role = replica").Error; err != nil {
		return err
	}
	for _, t := range ts {
		if err := db.Exec(fmt.Sprintf("truncate table %s", t)).Error; err != nil {
			return err
		}
	}
	if err := db.Exec("SET session_replication_role = DEFAULT").Error; err != nil {
		return err
	}
	return nil
}
