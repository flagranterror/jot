# Jot

Jot appends a timestamped entry to a daily journal file. I wrote because I love
the "Append to Journal" task in [Drafts](http://agiletortoise.com/drafts/) on
IOS, and wanted a work-alike for my desktop.

## Usage

```
Usage of ./jot_linux32:
  -datefmt string
    	Date Format - Changing the date format can change how often a file is rotated. (see golang time package) (default "Jan-2006")
  -dir string
    	Notes Directory (default "/home/user/Dropbox/Notes")
  -note string
    	Note content (default "Tick..")
  -pre string
    	Note filename prefix (default "jot-")
  -timefmt string
    	Timestamp Format (see golang time package) (default "## Jan 2 15:04:05")

```

By default creates or appends to `$HOME/Dropbox/Notes/jot-Jan-02-2006.txt`,
substituting in today's date.

Jot was designed to be called from a small bash/Zenity script, so it only
supports command line configuration. A bash (or batch for windows) script
is trivially easy to create. I've included a couple, and the Zenity wrapper
I'm using. I'm sure it would be equally simple to put something together for
Alfred on OS X.

## Building

* Install [Go](https://golang.org/)
* `go get github.com/easy-bot/jot`
* `go install github.com/easy-bot/jot`


## Dependencies

Jot doesn't do enough to require dependencies.
