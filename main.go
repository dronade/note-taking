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
	// check if table exists
	// if not, create table
	if _, err := os.Stat("sqlite-database.db"); os.IsNotExist(err) {
		file, err2 := os.Create("sqlite-database.db")
		if err2 != nil {
			log.Fatal(err2.Error())
		}
		file.Close()
	}
	createTable()
	ui.Main(setupUI)
	// insertNote("very important stuff!!!", "this, this, this and this")
	// insertNote("blah", "blh blah blah blah blah")
	// insertNote("ahem", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam accumsan lobortis arcu, rhoncus cursus augue malesuada in. Sed eget bibendum magna")
	// insertNote("to be deleted", "if this is stil in the table something's gone wrong :(")
	// deleteNote(sqliteDatabase)
	// displayNotes(sqliteDatabase)
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

	noteTitle := ui.NewEntry()
	noteContent := ui.NewMultilineEntry()
	entryForm.Append("Title", noteTitle, false)
	entryForm.Append("Enter Note:", noteContent, true)

	button := ui.NewButton("Create")

	vbox.Append(button, false)
	button.OnClicked(func(*ui.Button) {
		log.Println(noteTitle.Text())
		log.Println(noteContent.Text())
		go insertNote(noteTitle.Text(), noteContent.Text())
		noteTitle.SetText("")
		noteContent.SetText("")
	})

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

func createTable() {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer db.Close()
	createNoteTableSQL := `CREATE TABLE IF NOT EXISTS notes (
		"idNote" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"content" TEXT		
	  );`
	statement, err := db.Prepare(createNoteTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func insertNote(title string, content string) {
	log.Println("inserting notes")

	// check if string is empty
	if title == "" || content == "" {
		log.Println("title or content is empty")
		return
	}

	db, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer db.Close()
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

func displayNotes() {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer db.Close()
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

func deleteNote() {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer db.Close()
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
