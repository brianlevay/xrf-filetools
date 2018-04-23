package configureSpe

import ()

type Configuration struct {
	Threshold  float64 `json:"-"`
	GainMinKeV float64 `json:"-"`
	GainMaxKeV float64 `json:"-"`
}

func ReadConfig() *Configuration {
	var config = &Configuration{Threshold: 1000.0, GainMinKeV: 0.02000, GainMaxKeV: 0.02050}
	return config
}
