package main

import (
	"flag"
	"github.com/mablo/df-seeker/pkg/output"
	"github.com/mablo/df-seeker/pkg/seek"
)

func main () {
	options := seek.Options{
		Path: flag.String(
			"p",
			".",
			"Directory path."),
		Recursive: flag.Bool(
			"r",
			false,
			"Recursive"),
		OpenFilesLimitInPercent: flag.Uint(
			"rlimit",
			90,
			"Percent usage of soft ulimit."),
		SortParameter: flag.String(
			"sp",
			"size",
			"Sort parameter (hash, size)."),
		SortOrder: flag.String(
			"so",
			"asc",
			"Sort order (asc, desc)."),
	}
	flag.Parse()

	output.Show(seek.Execute(options))
}
