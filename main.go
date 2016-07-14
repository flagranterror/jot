package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"
)

func main() {

	home := os.Getenv("HOME")

	notepath := flag.String("dir", fmt.Sprintf("%v/Dropbox/Notes", home), "Notes Directory")
	prefix := flag.String("pre", "jot-", "Note filename prefix")
	suffix := flag.String("suffix", ".md", "Note filename prefix")
	note := flag.String("note", "Tick..", "Note content")
	datefmt := flag.String("datefmt", "Jan-2006", "Date Format - Changing the date format can change how often a file is rotated. (see golang time package)")
	timefmt := flag.String("timefmt", "## Jan 2 15:04:05", "Timestamp Format (see golang time package)")

	flag.Parse()

	t := time.Now()
	date := t.Format(*datefmt)
	timestamp := t.Format(*timefmt)

	notefile := path.Join(*notepath, fmt.Sprintf("%v%v%v", *prefix, date, *suffix))

	f, err := os.OpenFile(notefile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%v\n%v\n\n", timestamp, *note)); err != nil {
		panic(err)
	}

}
