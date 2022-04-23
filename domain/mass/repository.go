package mass

import (
	"github.com/gsousadev/ea_liturgy/infra/database"

	"gorm.io/gorm"
)

func storeMassRepository(m *Mass) *gorm.DB {
	return database.Db.Create(m)
}

func deleteMassRepository(id int) *gorm.DB {
	return database.Db.Delete(&Mass{}, id)
}
