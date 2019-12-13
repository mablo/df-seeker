package fs

import "io/ioutil"

func FetchFiles(path string, recursive bool) map[int64][]string {
	return fetch(path, recursive, map[int64][]string {})
}

func fetch(path string, recursive bool, files map[int64][]string) map[int64][]string {
	items, _ := ioutil.ReadDir(path)

	for _, item := range items {
		if item.IsDir()  {
			if recursive {
				files = fetch(path + "/" + item.Name(), true, files)
			}
		} else {
			files[item.Size()] = append(files[item.Size()], path + "/" + item.Name())
		}
	}

	return files
}
