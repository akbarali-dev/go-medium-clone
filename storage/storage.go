package storage

import (
	"mediumclone/storage/postgres"
	"mediumclone/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
}

func NewStoragePg(psqlConn *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUserStorage(psqlConn),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}
