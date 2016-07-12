#!/bin/bash

# Example of a bash wrapper that can be used to set command line options.
note="$@"

$GOPATH/bin/jot -note="$note"
