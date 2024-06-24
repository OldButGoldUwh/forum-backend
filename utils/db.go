// forum/utils/db.go

package utils

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func GetDB() *sql.DB {
	return db
}

func createTables() {
	userTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
		token TEXT NOT NULL
    );`

	postTable := `CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        user_id INTEGER NOT NULL,
		categories TEXT NOT NULL,
        likes INTEGER NOT NULL,
        dislikes INTEGER NOT NULL,
		comment_length INTEGER NOT NULL,
        created_at TEXT NOT NULL,
        updated_at TEXT NOT NULL
    );`

	commentTable := `CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER NOT NULL,
        content TEXT NOT NULL,
        user_id INTEGER NOT NULL,
		created_at TEXT NOT NULL
    );`

	likeTable := `CREATE TABLE IF NOT EXISTS likes (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER,
        user_id INTEGER NOT NULL,
		comment_id INTEGER,
		like BOOl NOT NULL,
		created_at TEXT NOT NULL
	);`

	_, err := db.Exec(userTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(postTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(commentTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(likeTable)
	if err != nil {
		log.Fatal(err)
	}
}
