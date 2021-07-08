package main

import (
	"github.com/scottski/di-test/pkg/externalDB1"
	"github.com/scottski/di-test/pkg/service"
)

func main() {

	msg := `"{"stuff": "abc"}"`

	db := externalDB1.NewDB()

	msgService := service.NewMessageService(&db)
	msgService.RecieveMessage(msg)
}
