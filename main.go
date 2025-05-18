package main

import (
	"crud-go/config"
	"crud-go/migrations"
	"crud-go/routers"
)

func main() {

	config.ConnectDatabase()
	db := config.DB
	migrations.MigrateAll(db)

	r := routers.SetupRouter(db)
	r.Run(":8080")
}
