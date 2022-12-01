package puzzleinput

import (
	"embed"
	"fmt"
	"log"
)

//go:embed files/*
var inputFiles embed.FS

// Day loads the input for a given day. panics on any errors
func Day(daynum int) []byte {
	res, err := inputFiles.ReadFile(fmt.Sprintf("files/day%d.txt", daynum))
	if err != nil {
		log.Fatalf("couldn't get input for day %d: %v", daynum, err)
	}
	return res
}
