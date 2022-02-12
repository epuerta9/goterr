package main

import (
	"encoding/json"
	"log"
	"os"
	"text/template"
)

func checkPipe() bool {
	//check to make sure pipe exists
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	} else {
		return true
	}

}

var (
	tf ResourceChanges
)

func main() {
	if !checkPipe() {
		log.Fatal("No input from pipe detected")
	}

	//after checking pipe, going to unmarshal terraform data
	if err := json.NewDecoder(os.Stdin).Decode(&tf); err != nil {
		log.Fatal("unable to unmarshal json")
	}

	tpl := template.Must(template.ParseFiles("body.template"))

	tpl.Execute(os.Stdout, tf)

}
