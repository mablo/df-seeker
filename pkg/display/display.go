package display

import (
	"fmt"
	"github.com/mablo/df-seeker/pkg/duplicate"
)

func Display(duplicates []duplicate.Duplicate) {
	for _, dup := range duplicates {
		fmt.Printf("size: %s, hash: %s\n", humanSize(dup.Size), dup.Sum)
		for _, file := range dup.Files {
			fmt.Printf("- %s\n", file)
		}
		fmt.Print("\n")
	}
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