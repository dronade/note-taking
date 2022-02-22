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
	file, err := os.Create("note2.txt")
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
	f, err := os.Create("note2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("write to note")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
}

func editNote() {

}

func removeNote() {

}

func main() {
	//ask user what they want to do
	createNote()
	readNote()
	writetoNote()
}
