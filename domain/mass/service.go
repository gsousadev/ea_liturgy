package mass

import (
	"encoding/json"
)

func storeMassService(data *MassRequest) (Mass, error) {
	out, _ := json.Marshal(data.CelebratoryBody)

	m := Mass{
		Description:     data.Description,
		Date:            data.Date,
		Hour:            data.Hour,
		Location:        data.Location,
		CelebratoryBody: string(out),
	}

	result := storeMassRepository(&m)

	if result.Error != nil {
		return m, result.Error
	}

	return m, nil
}
