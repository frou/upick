Purpose
=======

`upick` picks and launches a file of its own choosing. Either:

* A file in the current directory.
* A file in the current directory or any nested directories (with the `-r` flag).

The file will launch with the default application registered for its extension.

OS X, Windows and other Unixes (where the `xdg-open` command is available) are supported.

Treatment of Symbolic Links
---------------------------

* Symlinked files are eligible to be picked.
* Symlinked directory contents are not eligible to be picked because that would allow cycles.

Installation
============

    $ go get github.com/frou/upick
    $ # Command is installed in $GOPATH/bin
