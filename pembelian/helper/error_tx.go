package helper

import (
	"database/sql"
	"eh_teh_mewa/helperMain"
)

func ErrorTx(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		helperMain.PanicIfError(errorRollback)
	} else {
		errorCommit := tx.Commit()
		helperMain.PanicIfError(errorCommit)
	}
}
