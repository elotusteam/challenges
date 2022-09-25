package migration

import (
	. "TEST/quanngo/hackathon/migrator"
)

func addUserMigration(mg *Migrator) {
	role := Table{
		Name: "user",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "user_name", Type: DB_NVarchar, Length: 200, Nullable: false},
			{Name: "password", Type: DB_NVarchar, Length: 200, Nullable: false},
			{Name: "create_at", Type: DB_DateTime, Nullable: false, Default: DB_NowTimeZoneUTC},
		},
		Indices: []*Index{},
	}

	// create table
	mg.AddMigration("create user table", NewAddTableMigration(role))
}
