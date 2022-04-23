package mass

import (
	"github.com/gsousadev/ea_liturgy/infra/database"

	"gorm.io/gorm"
)

type Mass struct {
	gorm.Model
	Description     string
	Date            string
	Hour            string
	Location        string
	CelebratoryBody string
}

type Reader struct {
	gorm.Model
	Name  string
	Phone string
}

type Priest struct {
	gorm.Model
	Name  string
	Phone string
}

type MinisterOfTheEucharist struct {
	gorm.Model
	Name  string
	Phone string
}

type MusicMinister struct {
	gorm.Model
	Name  string
	Phone string
}

type Acolyte struct {
	gorm.Model
	Name  string
	Phone string
}

type AltarBoy struct {
	gorm.Model
	Name  string
	Phone string
}

type Ceremonial struct {
	gorm.Model
	Name  string
	Phone string
}

func MigrateStructs() {
	database.Db.AutoMigrate(
		Mass{},
		Reader{},
		Priest{},
		MinisterOfTheEucharist{},
		MusicMinister{},
		Acolyte{},
		AltarBoy{},
		Ceremonial{},
	)
}
