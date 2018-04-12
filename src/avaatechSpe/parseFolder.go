package avaatechSpe

import (
	"os"
	"strings"
)

func (spe *SPE) ParseFolder() error {
	var pathPts []string
	var pathLength int
	pathPts = strings.Split(spe.FilePath, string(os.PathSeparator))
	pathLength = len(pathPts)
	switch pathLength {
	case 0:
		{
			spe.Folder = "Root"
		}
	case 1:
		{
			spe.Folder = "Root"
		}
	case 2:
		{
			if strings.Contains(pathPts[pathLength-2], "Run") {
				spe.Folder = "Root"
			} else {
				spe.Folder = pathPts[pathLength-2]
			}
		}
	default:
		{
			if strings.Contains(pathPts[pathLength-2], "Run") {
				spe.Folder = pathPts[pathLength-3]
			} else {
				spe.Folder = pathPts[pathLength-2]
			}
		}
	}
	return nil
}
