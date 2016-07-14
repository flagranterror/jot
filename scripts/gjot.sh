#!/bin/bash
# Zenity gui for adding a journal entry. This is how I get stuff into my journal.

ZENITY=/usr/bin/zenity
JOT=~/bin/jot
note=$($ZENITY --entry --text="Line to log" --entry-text="Remember this moment." 2> /dev/null)

$JOT -note="$note"
