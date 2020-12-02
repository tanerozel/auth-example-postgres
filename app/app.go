package app

import (
	"encoding/gob"
	"log"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	err := godotenv.Load("env")
	if err != nil {
		log.Print(err.Error())
		return err
	}

	Store = sessions.NewFilesystemStore("", []byte("something-very-secret"))
	gob.Register(map[string]interface{}{})
	return nil
}
