package output

import (
	"fmt"
	"github.com/mablo/df-seeker/pkg/seek"
)

func Show(duplicates []seek.Duplicate) {
	for _, dup := range duplicates {
		fmt.Printf("size: %s, hash: %s\n", convertForHuman(dup.Size), dup.Hash)
		for _, file := range dup.Files {
			fmt.Printf("- %s\n", file.Path)
		}
		fmt.Print("\n")
	}
}