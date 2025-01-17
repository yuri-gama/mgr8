package cmd

import (
	"fmt"
	"os"

	"github.com/kenji-yamane/mgr8/domain"
)

type generate struct{}

func (g *generate) execute(args []string, databaseURL string, migrationsDir string, driver domain.Driver) error {
	filePath := args[0]
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read from file with path: %s", err)
	}

	_, err = driver.ParseMigration(string(content))
	if err != nil {
		return err
	}

	return nil
}
