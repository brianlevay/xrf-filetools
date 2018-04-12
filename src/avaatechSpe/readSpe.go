package avaatechSpe

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadSPE(path string, info os.FileInfo, open bool) (*SPE, error) {
	var errName, errCont error
	spe := new(SPE)
	spe.FilePath = path
	spe.FileName = info.Name()
	spe.Initialize()
	errName = spe.ParseFileName()
	if errName != nil {
		return nil, errName
	}
	if open == true {
		errCont = spe.ParseFileContents()
		if errCont != nil {
			return nil, errCont
		}
		spe.Opened = true
	}
	return spe
}
