package mass

type CelebratoryBodyRequest struct {
	Celebrant               PriestRequest                   `json:"celebrant" binding:"required"`
	AdditionalCelebrants    []PriestRequest                 `json:"additional_celebrants"`
	FirstReader             ReaderRequest                   `json:"first_reader" binding:"required"`
	SecondReader            ReaderRequest                   `json:"second_reader" binding:"required"`
	AssemblyPrayerReader    ReaderRequest                   `json:"assembly_prayer_reader" binding:"required"`
	MinistersOfTheEucharist []MinisterOfTheEucharistRequest `json:"ministers_of_the_eucharist" binding:"required"`
	MusicMinisters          []MusicMinisterRequest          `json:"music_ministers" binding:"required"`
	Acolytes                []AcolyteRequest                `json:"acolytes" binding:"required"`
	AltarBoys               []AltarBoyRequest               `json:"altar_boys" binding:"required"`
	Ceremonials             []CeremonialRequest             `json:"ceremonial" binding:"required"`
}

type MassRequest struct {
	Description     string                 `json:"description" binding:"required"`
	Date            string                 `json:"date" binding:"required"`
	Hour            string                 `json:"hour" binding:"required"`
	Location        string                 `json:"location" binding:"required"`
	CelebratoryBody CelebratoryBodyRequest `json:"celebratory_body" binding:"required"`
}

type ReaderRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type PriestRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type MinisterOfTheEucharistRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type MusicMinisterRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type AcolyteRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type AltarBoyRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type CeremonialRequest struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
