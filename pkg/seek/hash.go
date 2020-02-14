package seek

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/mablo/df-seeker/pkg/fs"
	"io"
	"os"
	"sync"
)

func calculateHash(file fs.File, wg *sync.WaitGroup, channel chan map[string]string, limit chan struct{}) {
	limit <- struct{}{}
	defer func() { <-limit }()
	defer wg.Done()

	fh, err := os.Open(file.Path)

	if err == nil {
		h := sha1.New()
		_, errr := io.Copy(h, fh)
		fh.Close()

		if errr == nil {
			hash := make(map[string]string)
			hash[file.Path] = hex.EncodeToString(h.Sum(nil))

			channel <- hash
		}
	} else {
		panic("I/O error - " + err.Error())
	}
}