package readAvaatechSpe

import (
	"path/filepath"
)

func ReadSPE(spePath string, open bool) (*SPE, error) {
	spe := new(SPE)
	spe.Initialize()
	spe.FilePath = spePath
	_, speName := filepath.Split(spePath)
	spe.FileName = speName
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
