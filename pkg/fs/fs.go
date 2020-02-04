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

func FetchFiles(path string, recursive bool) map[int64][]string {
	return fetch(strings.TrimRight(path, "/"), recursive, map[int64][]string {})
}

func fetch(path string, recursive bool, files map[int64][]string) map[int64][]string {
	items, _ := ioutil.ReadDir(path)

	for _, item := range items {
		switch mode := item.Mode(); {
		case mode.IsDir():
			if recursive {
				files = fetch(path+"/"+item.Name(), true, files)
			}
		case mode.IsRegular():
			files[item.Size()] = append(files[item.Size()], path + "/" + item.Name())
		}
	}

	return files
}
