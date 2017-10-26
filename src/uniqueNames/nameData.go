package uniqueNames

import ()

type UniqueNames struct {
	SourcePath string          `json:"-"`
	Error      string          `json:"Error"`
	Data       map[string]bool `json:"Data"`
}

func (unique *UniqueNames) Initialize(sourcePath string) {
	unique.SourcePath = sourcePath
	unique.Error = "none"
	unique.Data = map[string]bool{}
}
