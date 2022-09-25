package migration

import (
	. "TEST/quanngo/hackathon/migrator"
)

func AddMigrations(mg *Migrator) {
	addMigrationLogMigrations(mg)
	addUserMigration(mg)
	addFileMigration(mg)
}

func addMigrationLogMigrations(mg *Migrator) {
	migrationLogV1 := Table{
		Name: "migration_log",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "migration_id", Type: DB_NVarchar, Length: 255},
			{Name: "sql", Type: DB_Text},
			{Name: "success", Type: DB_Bool},
			{Name: "error", Type: DB_Text},
			{Name: "timestamp", Type: DB_DateTime},
		},
	}

	mg.AddMigration("create migration_log table", NewAddTableMigration(migrationLogV1))
}
