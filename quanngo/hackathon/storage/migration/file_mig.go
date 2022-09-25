package migration

import (
	. "TEST/quanngo/hackathon/migrator"
)

func addFileMigration(mg *Migrator) {
	role := Table{
		Name: "file",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "file_name", Type: DB_NVarchar, Length: 200, Nullable: false},
			{Name: "size", Type: DB_NVarchar, Length: 200, Nullable: false},
			{Name: "type", Type: DB_NVarchar, Length: 200, Nullable: false},
			{Name: "create_at", Type: DB_DateTime, Nullable: false, Default: DB_NowTimeZoneUTC},
		},
		Indices: []*Index{},
	}

	// create table
	mg.AddMigration("create file table", NewAddTableMigration(role))
}
