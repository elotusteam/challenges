package postgres

import (
	"fmt"

	"xorm.io/xorm"
)

type dbTransactionFunc func(sess *xorm.Session) error

func InTransaction(callback dbTransactionFunc) error {
	sess := X.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}

	if err := callback(sess); err != nil {
		if rollErr := sess.Rollback(); rollErr != nil {
			return fmt.Errorf("Failed to roll back transaction due to error: %s", rollErr)
		}
		return err
	}

	if err := sess.Commit(); err != nil {
		return err
	}

	return nil
}
