package database

import "golang.org/x/exp/maps"

type User struct {
	Name     string
	Password string

	Anime []Anime
}

type Anime struct {
	Name        string
	Description string
	Rating      float32
	Genre       []string
}

type Database interface {
	CreateUser(User)
	GetUser(username string) User
	CreateAnime(Anime)
	AllAnime() []Anime
	Anime(name string) Anime
}

type InMemory struct {
	users map[string]User
	anime map[string]Anime
}

func (db InMemory) CreateUser(user User) {
	db.users[user.Name] = user
}

func (db InMemory) GetUser(username string) User {
	user, exists := db.users[username]

	if !exists {
		println("User not found")
	}

	return user
}

func (db InMemory) CreateAnime(anime Anime) {
	db.anime[anime.Name] = anime
}

func (db InMemory) AllAnime() []Anime {
	return maps.Values(db.anime)
}

func (db InMemory) Anime(name string) Anime {
	anime, exists := db.anime[name]

	if !exists {
		println("Anime not found")
	}

	return anime
}

func NewInMemory() InMemory {
	db := InMemory{
		users: make(map[string]User),
		anime: make(map[string]Anime),
	}

	return db
}
