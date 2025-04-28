package storage

import "github.com/jmoiron/sqlx"

type StorageI interface {
}

type StoragePg struct{}

func NewStoragePg(psqlConn *sqlx.DB) *StorageI {
	return &StoragePg{}
}
