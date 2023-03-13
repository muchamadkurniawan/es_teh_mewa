package helperMain

import (
	"database/sql"
)

func ErrorTx(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}

func CloseRow(row *sql.Rows) {
	row.Close()
}
