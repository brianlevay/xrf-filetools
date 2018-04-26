package processAvaatechSpe

import (
	spereader "readAvaatechSpe"
)

type Spectrum struct {
	SPE   *spereader.SPE   `json:"SPE"`
	Peaks []*Peak          `json:"-"`
	Lines map[string]*Peak `json:"Lines"`
}

type Peak struct {
	Channel   float64 `json:"Channel"`
	HeightAbs float64 `json:"HeightAbs"`
	HeightRel float64 `json:"HeightRel"`
}
