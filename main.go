package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)
//need to add date created

func main() {
	os.Remove("sqlite-database.db")

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer sqliteDatabase.Close()
	createTable(sqliteDatabase)
	insertNote(sqliteDatabase, "very important stuff!!!", "this, this, this and this")
	insertNote(sqliteDatabase, "blah", "blh blah blah blah blah")
	insertNote(sqliteDatabase, "ahem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam accumsan lobortis arcu, rhoncus cursus augue malesuada in. Sed eget bibendum magna")
	displayNotes(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createNoteTableSQL := `CREATE TABLE notes (
		"idNote" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"content" TEXT		
	  );`

	log.Println("create notes table")
	statement, err := db.Prepare(createNoteTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("notes table created")
}

func insertNote(db *sql.DB, title string, content string) {
	log.Println("inserting notes")
	insertNoteSQL := `INSERT INTO notes(title, content) VALUES (?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(title, content)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayNotes(db *sql.DB) {
	row, err := db.Query("SELECT * FROM notes ORDER BY title")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		var id int
		var title string
		var content string
		row.Scan(&id, &title, &content)
		log.Println("note: ", title, " ", content, " ",)
	}
}
