package avaatechSpe

import ()

func ReadSPE(path string, name string, open bool) (*SPE, error) {
	spe := new(SPE)
	spe.Initialize()
	spe.FilePath = path
	spe.FileName = name
	errFolder := spe.ParseFolder()
	if errFolder != nil {
		return nil, errFolder
	}
	errName := spe.ParseFileName()
	if errName != nil {
		return nil, errName
	}
	if open == true {
		errCont := spe.ParseFileContents()
		if errCont != nil {
			return nil, errCont
		}
		spe.Opened = true
	}
	return spe, nil
}
