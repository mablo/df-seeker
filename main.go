package main

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)

var Files = map[int64][]string {}

type Duplicate struct {
	Size int64
	Sum string
	Files []string
}

var Duplicates []Duplicate

func main() {
	wordPtr := flag.String("path", ".", "Directory path.")
	flag.Parse()

	buildFilesMap(*wordPtr)

	var wg sync.WaitGroup
	hashesChannel := make(chan map[string]string)

	for _, list := range Files {
		if len(list) > 1 {
			for _, file := range list {
				wg.Add(1)
				go calculateHash(file, &wg, hashesChannel)
			}
		}
	}

	go func(messages chan map[string]string) {
		wg.Wait()
		close(messages)
	}(hashesChannel)

	a := getValues(hashesChannel)

	for size, list := range Files {
		if len(list) > 1 {
			hashes := map[string][]string {}

			for _, file := range list {
				hash := a[file]
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

	for _, dup := range Duplicates {
		fmt.Printf("size: %s, hash: %s\n", humanSize(dup.Size), dup.Sum)
		for _, file := range dup.Files {
			fmt.Printf("- %s\n", file)
		}
		fmt.Print("\n")
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

func calculateHash(file string, wg *sync.WaitGroup, channel chan map[string]string) {
	defer wg.Done()
	fh, _ := os.Open(file)

	h := sha1.New()
	io.Copy(h, fh)
	fh.Close()

	hash := make(map[string]string)

	hash[file] = hex.EncodeToString(h.Sum(nil))

	channel <- hash
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

func getValues(channel chan map[string]string) map[string]string {
	a := map[string]string {}
	for ret := range channel {
		for x, y := range ret {
			a[x] = y
		}
	}

	return a
}
