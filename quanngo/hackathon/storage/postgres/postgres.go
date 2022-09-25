package postgres

import (
	"fmt"
	"time"

	"TEST/quanngo/hackathon/migrator"
	migration "TEST/quanngo/hackathon/storage/migration"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

type Postgres struct {
}

var (
	X       *xorm.Engine
	Dialect migrator.Dialect
)

// New ....
func New(host, port, user, password, dbName string) (*Postgres, error) {
	var err error

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable binary_parameters=yes",
		host, port, user, password, dbName)

	orm, err := xorm.NewEngine("postgres", connect)
	if err != nil {
		return nil, err
	}
	orm.SetTZDatabase(time.UTC)

	X = orm
	Dialect = migrator.NewDialect(orm)

	migrator := migrator.NewMigrator(X)
	migration.AddMigrations(migrator)

	if err := migrator.Start(); err != nil {
		return nil, fmt.Errorf("Migration failed err: %v", err)
	}

	if err = X.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{}, nil
}

// Ping ...
func (p *Postgres) Ping() error {
	if err := X.Ping(); err != nil {
		return err
	}
	return nil
}

func (p *Postgres) Close() error {
	return X.Close()
}
