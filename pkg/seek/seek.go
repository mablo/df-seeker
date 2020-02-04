package seek

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/mablo/df-seeker/pkg/display"
	"github.com/mablo/df-seeker/pkg/duplicate"
	"github.com/mablo/df-seeker/pkg/fs"
	"io"
	"os"
	"sync"
	"time"
)

type Options struct {
	Path *string
	Recursive *bool
}

func Execute(options Options) []duplicate.Duplicate {
	filesList := fs.FetchFiles(*options.Path, *options.Recursive)

	//flatFilesList := fs.FetchFilesFlat(*options.Path, *options.Recursive)

	var wg sync.WaitGroup
	hashesChannel := make(chan map[string]string)

	for _, list := range filesList {
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

	var duplicates []duplicate.Duplicate

	for size, list := range filesList {
		if len(list) > 1 {
			hashes := map[string][]string {}

			for _, file := range list {
				hash := a[file]
				hashes[hash] = append(hashes[hash], file)
			}

			for hash, files := range hashes {
				if len(files) > 1 {
					duplicates = append(duplicates, duplicate.Duplicate{
						Size:  size,
						Sum:   hash,
						Files: files,
					})
				}
			}
		}
	}

	return duplicates
}

func ExecuteAndDisplay(options Options) {
	display.Display(Execute(options))
}

func calculateHash(file string, wg *sync.WaitGroup, channel chan map[string]string) {
	fh, err1 := os.Open(file)

	if err1 != nil  {
		print(file, " - czekam: ", "\n")
		time.Sleep(1 * time.Second)
		calculateHash(file, wg, channel, )
		return
	}

	h := sha1.New()
	_, err := io.Copy(h, fh)
	fh.Close()

	//print (" \t", file, "\t", err1 == nil, "\t", err == nil, "\n")

	if err == nil {
		hash := make(map[string]string)
		hash[file] = hex.EncodeToString(h.Sum(nil))

		//print(hash[file], "\t", file, "\n")

		channel <- hash
	}

	wg.Done()
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
