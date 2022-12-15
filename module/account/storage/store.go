package accountstorage

import (
	"gorm.io/gorm"
)

type SqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *SqlStore {
	return &SqlStore{db: db}
}
