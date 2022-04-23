package main

import (
	"github.com/gsousadev/ea_liturgy/domain/mass"
	"github.com/gsousadev/ea_liturgy/infra/database"
	"github.com/gsousadev/ea_liturgy/infra/router"
)

func main() {
	database.InitDB()
	mass.MigrateStructs()
	router.InitServer()
}
