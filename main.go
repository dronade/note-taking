package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	_ "github.com/mattn/go-sqlite3"
)

// need to add date created
// allow user to add title and content
// allow deletion from table
// allow user to edit title and content
// allow a user to delete
// decide how to go about implementing the frontend

func main() {
	ui.Main(setupUI)
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
	insertNote(sqliteDatabase, "to be deleted", "if this is stil in the table something's gone wrong :(")
	deleteNote(sqliteDatabase)
	displayNotes(sqliteDatabase)
}

var mainwin *ui.Window

func makeCreateNotePage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	entryForm.Append("Title", ui.NewEntry(), false)
	entryForm.Append("Enter Note:", ui.NewMultilineEntry(), true)
	vbox.Append(ui.NewButton("Create", ), false)

	return vbox
}

func makeViewNotePage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	return hbox
}

func setupUI() {
	mainwin = ui.NewWindow("Note taking app", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("Create Note", makeCreateNotePage())
	tab.SetMargined(0, true)

	tab.Append("View Notes", makeViewNotePage())
	tab.SetMargined(1, true)

	mainwin.Show()
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
		log.Println("note: ", title, " ", content, " ")
	}
}

func deleteNote(db *sql.DB) {
	log.Println("deleting note")
	deleteNoteSQL := `DELETE FROM notes WHERE title="to be deleted"`
	statement, err := db.Prepare(deleteNoteSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
