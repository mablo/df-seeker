package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sync"
)

var fileList = []string {}
var hashes = map[string]string {}

const amount = 5

func main()  {
	wordPtr := flag.String("path", ".", "a string")

	flag.Parse()

	parseDirectory(*wordPtr)

	var wg sync.WaitGroup

	maxOffset := int(math.Ceil(float64(len(fileList))/float64(amount)))
	abc := make(chan map[string]string, maxOffset)
	wg.Add(maxOffset)

	for offset := 0; offset < maxOffset; offset++ {
		x := offset*amount
		y := (offset+1)*amount

		//fmt.Printf("%d, %d, %d\n", x, y, maxOffset)

		go func(files []string) {
			defer wg.Done()
			mappy := make(map[string]string)

			for _, file := range files {
				mappy[file] = getSize(file)
			}
			abc <- mappy
		}(fileList[x:y])
	}

	go func(messages chan map[string]string) {
		wg.Wait()
		close(messages)
	}(abc)

	a := map[string][]string {}

	for ret := range abc {
		//fmt.Print(ret)
		for x, y := range ret {
			a[y] = append(a[y], x)
		}
	}

	i := 0

	for _, ret  := range a {
		if len(ret) > 1 {
			hashes := map[string][]string {}

			for _, file := range ret {
				hash := getHash(file)
				hashes[hash] = append(hashes[hash], file)
			}

			for _, files := range hashes {
				if len(files) > 1 {
					fmt.Print("\n")
					for _, file := range files {
						fmt.Printf("%s\n", file)
					}

					fmt.Print("\n")
					i++
				}
			}
		}
	}

	//fmt.Print(a)



	//msg := <-abc
	//fmt.Print(msg)
	//wg.Wait()


	//item := make(map[string]string)
	//
	//for ret := range abc {
	//	for x, y := range ret {
	//		item[x] = y
	//	}
	//}
	//
	//b, err := json.MarshalIndent(item, "", "  ")
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Print(string(b))
}

func parseDirectory(path string) {
	files, _ := ioutil.ReadDir(path)

	for _, file := range files {
		if (file.IsDir()) {
			parseDirectory(path + "/" + file.Name())
		} else {
			fileList = append(fileList, path + "/" + file.Name())
		}
	}
}

func getHash(file string) string {
	fh, _ := os.Open(file)

	h := sha256.New()
	io.Copy(h, fh)

	fh.Close()

	return hex.EncodeToString(h.Sum(nil))
}

func getSize(file string) string {
	fi, err := os.Stat(file)

	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%8d", fi.Size())
}