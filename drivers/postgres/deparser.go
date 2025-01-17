package postgres

import (
	"fmt"
	"sort"

	"github.com/kenji-yamane/mgr8/domain"
)

type deparser struct { }

func inStringList(stringList []string, needle string) bool {
	isIn := false
	for _, s := range stringList {
		if needle == s {
			isIn = true
		}
	}
	return isIn
}

func hasSingleArg(datatype string) bool {
	singleArgTypes := []string{"char", "varchar", "bit", "varbit", "time", "timestamp"}
	if inStringList(singleArgTypes, datatype) {
		return true
	} else {
		return false
	}
}

func hasDoubleArg(datatype string) bool {
	doubleArgTypes := []string{"decimal", "numeric"}
	if inStringList(doubleArgTypes, datatype) {
		return true
	}	else {
		return false
	}
}

func (d *deparser) CreateTable(table *domain.Table) string {
 	statement := fmt.Sprintf("CREATE TABLE %s (\n", table.Name)

	columnKeys := []string{}
	for columnName, _ := range table.Columns {
		columnKeys = append(columnKeys, columnName)
	}
	sort.Strings(columnKeys)

	for _, key := range columnKeys {
		column := table.Columns[key]

		statement = statement + fmt.Sprintf("%s %s", key, column.Datatype)

		if hasSingleArg(column.Datatype) {
			statement = statement + fmt.Sprintf("(%d)", column.Parameters["size"])
		} else if hasDoubleArg(column.Datatype) {
			statement = statement + fmt.Sprintf("(%d,%d)", column.Parameters["precision"], column.Parameters["scale"])
		}

		if column.IsNotNull {
			statement = statement + fmt.Sprintf(" NOT NULL")
		}

		statement = statement + fmt.Sprintf(",\n")
	}

	statement = statement[0:len(statement) - 2]
	statement = statement + fmt.Sprintf("\n)")
	return statement
}

func (d *deparser) DropTable(tableName string) string {
	return fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
}

func (d *deparser) AddColumn() string {
	// TODO
	return ""
}

func (d *deparser) DropColumn(tableName, columnName string) string {
	return fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", tableName, columnName)
}
func (d *deparser) MakeColumnNotNull(tableName, columnName string) string {
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s SET NOT NULL", tableName, columnName)
}

func (d *deparser) UnmakeColumnNotNull(tableName, columnName string) string {
	return fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s DROP NOT NULL", tableName, columnName)
}
