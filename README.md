# Jot

Jot appends a timestamped entry to a daily journal file.


```
Usage of bin/jot:
  -datefmt string
    	Date Format (see golang time package) (default "Jan-02-2006")
  -dir string
    	Notes Directory (default "/home/user/Dropbox/Notes")
  -note string
    	Note content (default "Tick..")
  -pre string
    	Note filename prefix (default "jot-")
  -timefmt string
    	Timestamp Format (see golang time package) (default "15:04:05")
```

By default creates or appends to `$HOME/Dropbox/Notes/jot-Jan-02-2006.txt`,
substituting in today's date.

Jot was designed to be called from a small bash/Zenity script, so it only
supports command line configuration. A bash (or batch for windows) script is 
trivially easy to create. I've included a couple, and the Zenity wrapper I'm
using. I'm sure it would be equally simple to put something together for Alfred
on OS X.
