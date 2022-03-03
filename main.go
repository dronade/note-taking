package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readNote() {
	file, err := os.Open("note2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	b, err := ioutil.ReadAll(file)
	myString := string(b[:])
	fmt.Print("\n", myString)
}

func writetoNote() {
	file, err := os.Create("note2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := file.WriteString("write to note")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
}

func editNote() {
	file, err := os.OpenFile("note1.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString("\nfifth line"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//ask user what they want to do
	//get user input
	// fmt.Print("Enter what you want to do: ")
	// var userInput string
	// fmt.Scanf("%s", &userInput)

	// switch userInput {
	// case "create":
	// 	defer createNote()
	// case "read":
	// 	readNote()
	// case "write":
	// 	writetoNote()
	// case "edit":
	// 	editNote()
	// case "remove":
	// 	removeNote()
	// default:
	// 	fmt.Println("invalid command")
	// }
}
