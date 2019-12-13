package main

import (
	"flag"
	"github.com/mablo/df-seeker/pkg/seek"
)

func main () {
	options := seek.Options{
		Path: flag.String(
			"path",
			".",
			"Directory path."),
		Recursive: flag.Bool(
			"r",
			false,
			"Recursive"),
	}
	flag.Parse()

	seek.ExecuteAndDisplay(options)
}
