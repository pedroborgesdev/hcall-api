package database

import (
	"gorm.io/gorm"
)

func ExecuteInTransaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	// Inicia a transação
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
