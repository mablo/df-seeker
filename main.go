package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var Files = map[int64][]string {}

type Duplicate struct {
	Size int64
	Sum string
	Files []string
}

var Duplicates []Duplicate

func main()  {
	wordPtr := flag.String("path", ".", "Directory path.")

	flag.Parse()

	buildFilesMap(*wordPtr)

	for size, list := range Files {
		if len(list) > 1 {
			hashes := map[string][]string {}

			for _, file := range list {
				hash := calculateHash(file)
				hashes[hash] = append(hashes[hash], file)
			}

			for hash, files := range hashes {
				if len(files) > 1 {
					Duplicates = append(Duplicates, Duplicate{
						Size:  size,
						Sum:   hash,
						Files: files,
					})
				}
			}
		}
	}
}

func buildFilesMap(path string) {
	items, _ := ioutil.ReadDir(path)

	for _, item := range items {
		if item.IsDir() {
			buildFilesMap(path + "/" + item.Name())
		} else {
			Files[item.Size()] = append(Files[item.Size()], path + "/" + item.Name())
		}
	}
}

func calculateHash(file string) string {
	fh, _ := os.Open(file)

	h := sha256.New()
	io.Copy(h, fh)
	fh.Close()

	return hex.EncodeToString(h.Sum(nil))
}

func humanSize(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}