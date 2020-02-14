package seek

import (
	"github.com/mablo/df-seeker/pkg/fs"
	"sync"
)
func Execute(options Options) []Duplicate {
	files := fs.FilterBySize(fs.GroupBySize(fs.FetchFilesFlat(*options.Path, *options.Recursive)))

	var wg sync.WaitGroup
	hashesChannel := make(chan map[string]string)
	limit := make(chan struct{}, getRlimit(*options.OpenFilesLimitInPercent))

	for _, file := range files {
		wg.Add(1)
		go calculateHash(file, &wg, hashesChannel, limit)
	}

	go func(messages chan map[string]string) {
		wg.Wait()
		close(messages)
		close(limit)
	}(hashesChannel)

	return sortByOptions(*options.SortParameter, *options.SortOrder, updateHashes(getValues(hashesChannel), files))
}

func getValues(channel chan map[string]string) map[string]string {
	a := map[string]string {}
	for ret := range channel {
		for x, y := range ret {
			a[x] = y
		}
	}

	return a
}

func updateHashes(hashes map[string]string, files []fs.File) []Duplicate {
	tmpFiles := map[string][]fs.File{}
	for _, file := range files {
		file.Hash = hashes[file.Path]
		tmpFiles[file.Hash] = append(tmpFiles[file.Hash], file)
	}

	var duplicates []Duplicate

	for _, files := range tmpFiles {
		if len(files) > 1 {
			var s []fs.File = files[:1]

			duplicates = append(duplicates, Duplicate{
				Size:  s[0].Size,
				Hash:  s[0].Hash,
				Files: files,
			})
		}
	}

	return duplicates
}