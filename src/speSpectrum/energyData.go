package speSpectrum

type LinePair struct {
	Name   string
	Energy float64
}

// Source for Energies: Kaye and Laby, Table of Physical and Chemical Constants //
var lineList = []LinePair{
	LinePair{Name: "Al_Ka", Energy: 1.487},
	LinePair{Name: "Si_Ka", Energy: 1.740},
	LinePair{Name: "Rh_La", Energy: 2.698},
	LinePair{Name: "K_Ka", Energy: 3.314},
	LinePair{Name: "Ca_Ka", Energy: 3.692},
	LinePair{Name: "Ca_Kb", Energy: 4.013},
	LinePair{Name: "Fe_Ka", Energy: 6.404},
	LinePair{Name: "Fe_Kb", Energy: 7.058},
}

var lineMap = map[string]float64{
	"Al_Ka": 1.487,
	"Si_Ka": 1.740,
	"Rh_La": 2.698,
	"K_Ka":  3.314,
	"Ca_Ka": 3.692,
	"Ca_Kb": 4.013,
	"Fe_Ka": 6.404,
	"Fe_Kb": 7.058,
}
