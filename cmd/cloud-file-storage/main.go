package main

import (
	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/config"
)

func main() {
	// -- users --
	// todo connect to db postgres
	// todo create db + migrations
	// todo user model
	// todo auth/register
	// todo connect to redis
	// todo auth + session middleware
	// todo test auth

	// -- domain logic --
	// todo create domain models(Folder/File)
	// todo create domain logic(create/delete/rename/move object)
	// todo connect to object storage
	// todo operation and search service
	// todo test service and domain logic

	// -- http handle --
	// todo include router - chi
	// todo create endpoint api handler
	// todo open api spec

	// todo integrate frontend

	// -- deploy --
	// todo create docker image for go app
	// todo docker compose with redis, go, postgres and s3
	// todo deploy to server

	// -- other --
	// todo include logger - slog

	config.MustLoad()
}
