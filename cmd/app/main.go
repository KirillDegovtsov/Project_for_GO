package main

import (
	"log"
	objectHTTP "project_university/api/http"
	"project_university/cmd/config"
	pkgObject "project_university/pkg/http"
	ramstorage "project_university/repositoty/ram_storage"
	"project_university/usecases/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	appFlags := config.ParseFlags()
	var cfg config.Config
	config.MustLoad(appFlags.ConfigPath, &cfg)

	objectNoteDB := ramstorage.NewNote()
	objectNote, err := service.NewNote(objectNoteDB)
	if err != nil {
		log.Fatal("Failed to start server: #{err}")
	}

	objectUserDB := ramstorage.NewUser()
	objectUser := service.NewUser(objectUserDB)

	objectProvider := service.NewProvider()
	cookieName := "session_id"
	maxLifetime := int64(60)

	objectManager := service.NewManager(objectProvider, cookieName, maxLifetime)

	objectHandlers := objectHTTP.NewHandler(objectNote, objectUser, objectManager)

	r := chi.NewRouter()
	objectHandlers.WithObjectHandlers(r)

	if err := pkgObject.CreateAndRunServer(r, cfg.HTTPServer.Address); err != nil {
		log.Fatal("Failed to start server: #{err}")
	}
}
