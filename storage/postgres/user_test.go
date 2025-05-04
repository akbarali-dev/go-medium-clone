package postgres_test

// import "fmt"

import (
	"context"
	"mediumclone/internal/db"
	"mediumclone/storage"
	"testing"

	"github.com/google/uuid"
	// "github.com/jmoiron/sqlx"
	// "mediumclone/storage/postgres"
	"mediumclone/storage/repo"
)

// func TestUserRepo_Create(t *testing.T) {
// 	// 1. Baza bilan ulanamiz
// 	psqlConn := db.ConnectToPostgres()

// 	// 2. Transaction ochamiz
// 	tx, err := psqlConn.Beginx()
// 	if err != nil {
// 		t.Fatalf("❌ failed to begin transaction: %v", err)
// 	}
// 	defer tx.Rollback() // 3. Testdan keyin tozalash

// 	// 4. Transaction bilan ishlaydigan repo yaratamiz
// 	userRepo := postgres.NewUserStorage(tx) // <- tx ni uzatamiz

// 	// 5. Test ma'lumotini kiritamiz
// 	id := uuid.NewString()
// 	user := &repo.User{
// 		ID:        id,
// 		FirstName: "Ali",
// 		LastName:  "Valiyev",
// 		Email:     "ali1211@example.com",
// 		Password:  "123456",
// 	}

// 	created, err := userRepo.Create(context.Background(), user)
// 	if err != nil {
// 		t.Fatalf("❌ error creating user: %v", err)
// 	}
// 	if created.CreatedAt == nil || created.ID != id {
// 		t.Error("invalid user data")
// 	}

// }

func TestUserRepo_Create(t *testing.T) {
	psqlConn := db.ConnectToPostgres()

	strg := storage.NewStoragePg(psqlConn)
	id := uuid.NewString()
	user := &repo.User{
		ID:        id,
		FirstName: "Ali",
		LastName:  "Valiyev",
		Email:     "ali111@example.com",
		Password:  "123456",
	}

	created, err := strg.User().Create(context.Background(), user)
	if err != nil {
		t.Fatal(err)
	}
	if created.CreatedAt == nil || created.ID != id {
		t.Error("invalid user data")
	}
}

// func nima() {
// 	fmt.Println("nima")
// }
