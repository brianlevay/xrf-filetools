package processAvaatechSpe

import (
	spereader "readAvaatechSpe"
)

type Spectrum struct {
	SPE        *spereader.SPE   `json:"SPE"`
	MaxChannel int              `json:"-"`
	Signal     []float64        `json:"-"`
	Peaks      []*Peak          `json:"-"`
	Lines      map[string]*Peak `json:"Lines"`
	Gain       float64          `json:"-"`
	Offset     float64          `json:"-"`
	R2         float64          `json:"-"`
}

type Peak struct {
	Channel float64 `json:"Channel"`
	Height  float64 `json:"Height"`
	Total   float64 `json:"Total"`
}
