package drivers

import (
	"fmt"

	"github.com/kenji-yamane/mgr8/drivers/postgres"
)

type Driver interface {
	Begin(url string) error
	Commit() error

	Execute(statements []string) error
	GetLatestMigration() (int, error)
	UpdateLatestMigration(int) error
	CreateBaseTable() error
	HasBaseTable() (bool, error)
}

func GetDriver(driverName string) (Driver, error) {
	switch driverName {
	case "postgres":
		return postgres.NewPostgresDriver(), nil
	}
	return nil, fmt.Errorf("inexistent driver %s", driverName)
}
