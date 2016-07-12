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

	t := time.Now()
	date := t.Format("Jan-02-2006")
	timestamp := t.Format("15:04:05")

	notepath := flag.String("dir", fmt.Sprintf("%v/Dropbox/Notes", home), "Notes Directory")
	prefix := flag.String("pre", "jot-", "Note filename prefix")
	note := flag.String("note", "Tick..", "Note content")
	flag.Parse()

	notefile := path.Join(*notepath, fmt.Sprintf("%v%v.txt", *prefix, date))

	f, err := os.OpenFile(notefile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%v:\n%v\n\n", timestamp, *note)); err != nil {
		panic(err)
	}

}
