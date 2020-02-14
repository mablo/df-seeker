package seek

import (
	"sort"
)

func sortByOptions(parameter string, order string, duplicates []Duplicate) []Duplicate {
	switch parameter {
		case "h":
		case "hash": {
			sort.Slice(duplicates, func(i, j int) bool {
				if order == "desc" {
					return duplicates[i].Hash > duplicates[j].Hash
				} else {
					return duplicates[i].Hash < duplicates[j].Hash
				}
			})
		}
		case "s":
		case "size": {
			sort.Slice(duplicates, func(i, j int) bool {
				if order == "desc" {
					return duplicates[i].Size > duplicates[j].Size
				} else {
					return duplicates[i].Size < duplicates[j].Size
				}
			})
		}
		default: {
			panic("Invalid sort parameter. Use --help for available sort options.")
		}
	}

	return duplicates
}