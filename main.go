package main

import (
	"log"

	"github.com/kacpi52/t1_users/process"
)

func main() {
	addedUsersCollection := process.GetAndCreateUserConcurrently()
	log.Printf(`%d users has been added.`,len(addedUsersCollection.Collection))
}