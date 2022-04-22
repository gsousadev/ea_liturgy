package main

import (
	"ea_liturgy/domain/mass"
	"ea_liturgy/infra/database"
	"ea_liturgy/infra/router"
)

func main() {
	database.InitDB()
	mass.MigrateStructs()
	router.InitServer()
}
