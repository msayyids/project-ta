package helper

import "gorm.io/gorm"

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		// If there is a panic, rollback the transaction
		if rollbackErr := tx.Rollback().Error; rollbackErr != nil {
			// Handle the rollback error
			panic(rollbackErr)
		}
		panic(err)
	} else {
		// If no error, commit the transaction
		if commitErr := tx.Commit().Error; commitErr != nil {
			// Handle the commit error
			panic(commitErr)
		}
	}
}
