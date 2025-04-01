package database

import (
	"gorm.io/gorm"
)

// ExecuteInTransaction executa uma função dentro de uma transação
// Se a função retornar um erro, a transação é revertida
// Se a função executar com sucesso, a transação é confirmada
func ExecuteInTransaction(db *gorm.DB, fn func(tx *gorm.DB) error) error {
	// Inicia a transação
	tx := db.Begin()

	// Garante que a transação seja revertida em caso de pânico
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Executa a função passada como parâmetro dentro da transação
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	// Confirma a transação
	return tx.Commit().Error
}
