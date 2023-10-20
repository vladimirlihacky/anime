package main

import (
	"anime/database"
	"anime/server"
)

func main() {
	db := database.NewInMemory()

	db.CreateUser(database.User{
		Name:     "Nikita",
		Password: "qwerty123",
		Anime: []database.Anime{
			{"Блич", "Неведомая хуйня", 5, []string{}},
		},
	})

	server := server.NewServer(db)
	server.Start(":3000")
}
