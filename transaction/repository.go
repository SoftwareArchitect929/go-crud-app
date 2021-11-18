package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Transaction, error)
	FindByID(ID int) (Transaction, error)
	Create(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	Delete(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Find(&transactions).Error

	return transactions, err
}

func (r *repository) FindByID(ID int) (Transaction, error) {
	var transactions Transaction
	err := r.db.Find(&transactions, ID).Error

	return transactions, err
}

func (r *repository) Create(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repository) Delete(transaction Transaction) (Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
