package postgresql

import (
	s "github.com/Shemistan/uzum_shop/internal/storage"
	"github.com/jmoiron/sqlx"
)

const tableName = "products"

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) s.IStorage {
	return &storage{db: db}
}
