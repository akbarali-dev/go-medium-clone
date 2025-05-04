package main

import (
	// "context"
	// "fmt"

	// "github.com/google/uuid"

	// "github.com/jmoiron/sqlx"
	// "log"

	// "fmt"

	"log"

	_ "github.com/lib/pq"

	// "mediumclone/config"
	// "mediumclone/config"
	"mediumclone/config"
	"mediumclone/internal/db"
	"mediumclone/server"
	"mediumclone/storage"
	// "mediumclone/storage/repo"
	// "mediumclone/storage/repo"
)

func main() {
	// cfg := config.Load(".")
	// psqlUrl := fmt.Sprintf(
	// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	cfg.Postgres.Host,
	// 	cfg.Postgres.Port,
	// 	cfg.Postgres.User,
	// 	cfg.Postgres.Password,
	// 	cfg.Postgres.Database,
	// )

	// psqlConn, err := sqlx.Open("postgres", psqlUrl)
	// if err != nil {
	// 	log.Fatalf("Failed to conect to postgresql: %v", err)
	// }
	// log.Println("Postgres Connection successfull 👌")

	psqlConn := db.ConnectToPostgres()
	strg := storage.NewStoragePg(psqlConn)

	router := server.NewServer(&server.Options{
		Strg: strg,
	})
	port := config.Load("").Port

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run to server [%v]: %v", port, err)
	}

	// TODO get user by ID
	// user, err := strg.User().Get(context.TODO(), "6eca80af-9fa0-41eb-8189-a1a8afdfa95a")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)
	// TODO update user
	// err := strg.User().Update(context.TODO(), &repo.UpdateUser{
	// 	FirstName: "kimdir",
	// 	LastName:  "nimadir",
	// 	ID:        "6eca80af-9fa0-41eb-8189-a1a8afdfa95a",
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Todo delete user by id
	// err := strg.User().Delete(context.Background(), "6eca80af-9fa0-41eb-8189-a1a8afdfa95a")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(user)

	// id, err := uuid.NewRandom()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// user, err := strg.User().Create(context.TODO(), &repo.User{
	// 	ID:        id.String(),
	// 	FirstName: "Akbarali1",
	// 	LastName:  "ASqaraliy1ev",
	// 	Email:     "alwesisds3sw12@gmail.com",
	// 	Password:  "12345678",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)
}
