package uniqueNames

import (
	"fmt"
)

// Highest level function //
func (unique *UniqueNames) WriteToCSV(outPath string, outName string) error {
	for key, value := range unique.Data {
		fmt.Println("Key:", key, "Value:", value)
	}
	return nil
}
