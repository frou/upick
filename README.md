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

Example Usage
=============

    project_ideas $ upick
    Picked: webscale_database.txt
    project_ideas $ # File is opened with Sublime Text.

    presentations $ upick
    Nothing to pick from.
    presentations $ upick -r
    Picked: Rich Hickey/2012 - The Value Of Values.mp4
    presentations $ # File is opened with QuickTime.
