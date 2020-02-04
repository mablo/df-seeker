package main

import (
	"flag"
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
	}
	flag.Parse()

	seek.ExecuteAndDisplay(options)
}
