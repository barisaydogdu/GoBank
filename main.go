package main

import (
	"fmt"
	"log"
)

func main() {
	//Veritanbanı interface implement
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)

	//Create Tables
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
