package seek

import "github.com/mablo/df-seeker/pkg/fs"

type Duplicate struct {
	Size  int64
	Hash  string
	Files []fs.File
}
