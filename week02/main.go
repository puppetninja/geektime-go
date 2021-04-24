package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
)

func GetExistUser() error {
	dbENR := sql.ErrNoRows
	return errors.Wrap(dbENR, "Couldn't find exist user")
}

func GetPaidUsers() error {
	dbENR := sql.ErrNoRows
	return errors.Wrap(dbENR, "Get User Hit Error")
}

func main() {
	err := GetPaidUsers()
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("No paid users found so far!")
	}

	err = GetExistUser()
	if errors.Is(err, sql.ErrNoRows) {
		log.Fatalf("Couldn't find target user\nstack trace: \n%+v\n", err)
	}
}
