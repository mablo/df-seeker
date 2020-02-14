package fs

import (
	"io/ioutil"
	"strings"
)

func FetchFilesFlat(path string, recursive bool) []File {
	return fetchFlat(strings.TrimRight(path, "/"), recursive, []File {})
}

func fetchFlat(path string, recursive bool, files []File) []File {
	items, _ := ioutil.ReadDir(path)

	for _, item := range items {
		switch mode := item.Mode(); {
		case mode.IsDir():
			if recursive {
				files = fetchFlat(path + "/" + item.Name(), true, files)
			}
		case mode.IsRegular():
			files = append(files, File{
				item.Name(),
				path + "/" + item.Name(),
				item.Size(),
				"",
			})
		}
	}

	return files
}

func GroupBySize(files []File) map[int64][]File {
	group := map[int64][]File{}

	for _, file := range files {
		group[file.Size] = append(group[file.Size], file)
	}

	return group
}

func FilterBySize(filesBySize map[int64][]File) []File {
	var filtered []File

	for _, files := range filesBySize {
		if len(files) > 1 {
			filtered = append(filtered, files...)
		}
	}

	return filtered
}
