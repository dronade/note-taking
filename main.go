package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func createNote() {
	// be able to create note name
	// multiple line notes?
	file, err := os.Create("note3.txt")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note: ")
	text, _ := reader.ReadString('\n')

	input, err := file.WriteString(text)
	if err != nil {
		fmt.Println(err)
	}

	print(input)
	file.Close()
}

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

func removeNote() {
	err := os.Remove("note3.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	//ask user what they want to do
	
	//removeNote()
	createNote()
	readNote()
	writetoNote()
	editNote()
}
