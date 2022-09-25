package migrator

import (
	"fmt"
	"time"

	"xorm.io/xorm"
)

type Migrator struct {
	x          *xorm.Engine
	Dialect    Dialect
	migrations []Migration
}

type MigrationLog struct {
	Id          int64
	MigrationID string `xorm:"migration_id"`
	SQL         string `xorm:"sql"`
	Success     bool
	Error       string
	Timestamp   time.Time
}

func NewMigrator(engine *xorm.Engine) *Migrator {
	mg := &Migrator{}
	mg.x = engine
	mg.migrations = make([]Migration, 0)
	mg.Dialect = NewDialect(mg.x)
	return mg
}

func (mg *Migrator) MigrationsCount() int {
	return len(mg.migrations)
}

func (mg *Migrator) AddMigration(id string, m Migration) {
	m.SetId(id)
	mg.migrations = append(mg.migrations, m)
}

func (mg *Migrator) GetMigrationLog() (map[string]MigrationLog, error) {
	logMap := make(map[string]MigrationLog)
	logItems := make([]MigrationLog, 0)

	exists, err := mg.x.IsTableExist(new(MigrationLog))
	if err != nil {
		return nil, err
	}

	if !exists {
		return logMap, nil
	}

	if err = mg.x.Find(&logItems); err != nil {
		return nil, err
	}

	for _, logItem := range logItems {
		if !logItem.Success {
			continue
		}
		logMap[logItem.MigrationID] = logItem
	}

	return logMap, nil
}

func (mg *Migrator) Start() error {

	logMap, err := mg.GetMigrationLog()
	if err != nil {
		return err
	}

	migrationsPerformed := 0
	migrationsSkipped := 0
	for _, m := range mg.migrations {
		m := m
		_, exists := logMap[m.Id()]
		if exists {
			migrationsSkipped++
			continue
		}

		sql := m.SQL(mg.Dialect)

		record := MigrationLog{
			MigrationID: m.Id(),
			SQL:         sql,
			Timestamp:   time.Now(),
		}

		err := mg.inTransaction(func(sess *xorm.Session) error {
			err := mg.exec(m, sess)
			if err != nil {
				record.Error = err.Error()
				if _, err := sess.Insert(&record); err != nil {
					return err
				}
				return err
			}
			record.Success = true
			_, err = sess.Insert(&record)
			if err == nil {
				migrationsPerformed++
			}
			return err
		})
		if err != nil {
			return fmt.Errorf("%v: %w", "migration failed", err)
		}
	}

	return mg.x.Sync2()
}

func (mg *Migrator) exec(m Migration, sess *xorm.Session) error {

	condition := m.GetCondition()
	if condition != nil {
		sql, args := condition.Sql(mg.Dialect)

		if sql != "" {
			results, err := sess.SQL(sql, args...).Query()
			if err != nil {
				return err
			}

			if !condition.IsFulfilled(results) {
				return nil
			}
		}
	}

	var err error
	if codeMigration, ok := m.(CodeMigration); ok {
		err = codeMigration.Exec(sess, mg)
	} else {
		sql := m.SQL(mg.Dialect)
		_, err = sess.Exec(sql)
	}

	if err != nil {
		return err
	}

	return nil
}

type dbTransactionFunc func(sess *xorm.Session) error

func (mg *Migrator) inTransaction(callback dbTransactionFunc) error {
	sess := mg.x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if err := callback(sess); err != nil {
		if rollErr := sess.Rollback(); rollErr != nil {
			return fmt.Errorf("failed to roll back transaction due to error: %s", rollErr)
		}

		return err
	}

	if err := sess.Commit(); err != nil {
		return err
	}

	return nil
}
