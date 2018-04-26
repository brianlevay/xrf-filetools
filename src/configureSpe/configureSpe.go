package configureSpe

import ()

type Configuration struct {
	UTCoffset  string  `json:"-"`
	StdsPath   string  `json:"-"`
	Threshold  float64 `json:"-"`
	GainMinKeV float64 `json:"-"`
	GainMaxKeV float64 `json:"-"`
}

func ReadConfig() *Configuration {
	var config = &Configuration{
		UTCoffset:  "-05:00",
		StdsPath:   `./_misc/testData/standards/`,
		Threshold:  1000.0,
		GainMinKeV: 0.02000,
		GainMaxKeV: 0.02050,
	}
	return config
}
