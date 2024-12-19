package repository

import "gorm.io/gorm"

type BaseRepository interface {
	WithPreloads(tx *gorm.DB, preloads map[string][]interface{}) *gorm.DB
	WithTransaction(fn func(tx *gorm.DB) error) error
	SingleTransaction() *gorm.DB
	BeginTransaction() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB)
}

type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db}
}

func (r *baseRepository) WithPreloads(tx *gorm.DB, preloads map[string][]interface{}) *gorm.DB {
	for query, args := range preloads {
		tx = tx.Preload(query, args...)
	}

	return tx
}

func (r *baseRepository) WithTransaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *baseRepository) SingleTransaction() *gorm.DB {
	return r.db
}

func (r *baseRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *baseRepository) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *baseRepository) Rollback(tx *gorm.DB) {
	tx.Rollback()
}
