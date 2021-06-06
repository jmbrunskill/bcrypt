# bcrypt #

A quick hack bcrypt for the commandline

It reads a password via standard in and returns a bcrypt hash using default settings to standard out.

Note: bcrypt removes leading and trailing quote characters e.g. " as these are often accidentally added on Windows platforms

__usage examples__

`pwgen -s -1 14 | bcrypt`

`echo "HI" | bcrypt`