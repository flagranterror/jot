package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func prepend(fname string, note string, timestamp string, newline bool) error {

	buf, _ := ioutil.ReadFile(fname)
  format := "%v%v\n\n%v"

	if newline {
		format = "%v\n%v\n\n%v"
	}

	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()

	if err != nil {
		return err
	}

	if _, err := f.WriteString(fmt.Sprintf(format, timestamp, note, string(buf))); err != nil {
		return err
	}

	return nil
}

func append(fname string, note string, timestamp string, newline bool) error {

	format := "%v%v\n\n"

	if newline {
		format = "%v\n%v\n\n"
	}

	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf(format, timestamp, note)); err != nil {
		return err
	}

	return nil
}

func main() {

	home := os.Getenv("HOME")

	pre := flag.Bool("prepend", false, "Add note to beginning of file.")
	notepath := flag.String("dir", fmt.Sprintf("%v/Dropbox/Notes", home), "Notes Directory")
	prefix := flag.String("pre", "jot-", "Note filename prefix")
	suffix := flag.String("suffix", ".md", "Note filename prefix")
	note := flag.String("note", "Tick..", "Note content")
	datefmt := flag.String("datefmt", "Jan-2006", "Date Format - Changing the date format can change how often a file is rotated. (see golang time package)")
	timefmt := flag.String("timefmt", "**Mon, Jan 2 15:04:05:** ", "Timestamp Format (see golang time package)")
  newline := flag.Bool("newline", false, "Print a newline after datetime stamp.")

	flag.Parse()

	t := time.Now()
	date := t.Format(*datefmt)
	timestamp := t.Format(*timefmt)

	notefile := path.Join(*notepath, fmt.Sprintf("%v%v%v", *prefix, date, *suffix))

	if *pre {
		if err := prepend(notefile, *note, timestamp, *newline); err != nil {
			panic(err)
		}

	} else {
		if err := append(notefile, *note, timestamp, *newline); err != nil {
			panic(err)
		}
	}

}
