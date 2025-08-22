package halpers

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		IfPanicError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		IfPanicError(errorCommit)
	}
}
