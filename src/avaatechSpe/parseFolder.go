package avaatechSpe

import (
	"errors"
	"os"
	"strings"
)

func (spe *SPE) ParseFolder() error {
	var pathPts []string
	var pathLength int
	pathPts = strings.Split(spe.FilePath, string(os.PathSeparator))
	pathLength = len(pathPts)
	if pathLength == 1 {
		spe.Folder = "Root"
	} else if pathLength == 2 {
		if strings.Contains(pathPts[pathLength-2], "Run") {
			spe.Folder = "Root"
		} else {
			spe.Folder = pathPts[pathLength-2]
		}
	} else if pathLength >= 3 {
		if strings.Contains(pathPts[pathLength-2], "Run") {
			spe.Folder = pathPts[pathLength-3]
		} else {
			spe.Folder = pathPts[pathLength-2]
		}
	} else {
		return errors.New("Unable to parse parent folder for SPE")
	}
	return nil
}
