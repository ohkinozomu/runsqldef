package sqldefwrapper

import (
	"fmt"
	"strconv"

	"github.com/sqldef/sqldef"
	"github.com/sqldef/sqldef/database"
	"github.com/sqldef/sqldef/database/mssql"
	"github.com/sqldef/sqldef/database/mysql"
	"github.com/sqldef/sqldef/database/postgres"
	"github.com/sqldef/sqldef/database/sqlite3"
	"github.com/sqldef/sqldef/parser"
	"github.com/sqldef/sqldef/schema"
)

// RunCore is the core function that can be used directly in Go projects without CGO
func RunCore(dbType, dbName, user, password, host, portStr, schemaFile string, enableDropTable bool) error {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("failed to convert port: %v", err)
	}

	var db database.Database
	var sqlParser database.Parser
	var genMode schema.GeneratorMode

	switch dbType {
	case "mysql":
		db, err = mysql.NewDatabase(database.Config{
			DbName:   dbName,
			User:     user,
			Password: password,
			Host:     host,
			Port:     port,
		})
		if err != nil {
			return fmt.Errorf("failed to create a MySQL database adapter: %v", err)
		}
		sqlParser = database.NewParser(parser.ParserModeMysql)
		genMode = schema.GeneratorModeMysql

	case "postgres":
		db, err = postgres.NewDatabase(database.Config{
			DbName:   dbName,
			User:     user,
			Password: password,
			Host:     host,
			Port:     port,
		})
		if err != nil {
			return fmt.Errorf("failed to create a PostgreSQL database adapter: %v", err)
		}
		sqlParser = database.NewParser(parser.ParserModePostgres)
		genMode = schema.GeneratorModePostgres

	case "sqlite3":
		db, err = sqlite3.NewDatabase(database.Config{
			DbName: dbName,
		})
		if err != nil {
			return fmt.Errorf("failed to create a SQLite3 database adapter: %v", err)
		}
		sqlParser = database.NewParser(parser.ParserModeSQLite3)
		genMode = schema.GeneratorModeSQLite3

	case "mssql":
		db, err = mssql.NewDatabase(database.Config{
			DbName:   dbName,
			User:     user,
			Password: password,
			Host:     host,
			Port:     port,
		})
		if err != nil {
			return fmt.Errorf("failed to create a MSSQL database adapter: %v", err)
		}
		sqlParser = database.NewParser(parser.ParserModeMssql)
		genMode = schema.GeneratorModeMssql

	default:
		return fmt.Errorf("unknown database type: %s", dbType)
	}

	desiredDDLs, err := sqldef.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}
	options := &sqldef.Options{
		DesiredDDLs:     desiredDDLs,
		EnableDropTable: enableDropTable,
	}
	sqldef.Run(genMode, db, sqlParser, options)
	return nil
}
