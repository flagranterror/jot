ECHO OFF
set /p note=Enter Note: 
REM This is just an example batch script one could use to set defaults. It might even work.
jot.exe -dir="%USERPROFILE%/Dropbox/Notes" -note="%note%"

ECHO ON
