package mass

import (
	"ea_liturgy/infra/database"

	"gorm.io/gorm"
)

func storeMassRepository(m *Mass) *gorm.DB {
	return database.Db.Create(m)
}
