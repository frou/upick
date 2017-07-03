package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/frou/stdext"
)

var (
	recursive = flag.Bool("r", false,
		"recursive: pick from subdirectories too")

	includeDotfiles = flag.Bool("d", false,
		"dotfiles: whether dotfiles are eligible to be picked")

	interactive = flag.Bool("i", false,
		"interactive: don't exit, pick again upon each line written to stdin")

	match = flag.String("m", "",
		"regular expression that whitelists eligible file basenames")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	stdext.Exit(run())
}

func run() error {
	lineReader := bufio.NewScanner(os.Stdin)

	matchRe, err := regexp.Compile(*match)
	if err != nil {
		return err
	}

	for {
		candidatePaths, err := findCandidates(".", matchRe)
		if err != nil {
			return err
		}
		pick, err := pick(candidatePaths)
		if err != nil {
			return err
		}
		fmt.Print("Picked:", pick)

		if err := stdext.Launch(pick); err != nil {
			return err
		}

		if *interactive {
			lineReader.Scan()
		} else {
			fmt.Println()
			break
		}
	}
	return lineReader.Err()
}

func findCandidates(
	startPath string,
	matchRe *regexp.Regexp) ([]string, error) {

	var candidates []string
	err := filepath.Walk(startPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			basename := filepath.Base(path)

			if !*includeDotfiles && len(basename) > 1 && basename[0] == '.' {
				if info.IsDir() {
					return filepath.SkipDir
				} else {
					return nil
				}
			}

			if info.IsDir() {
				// path is a directory according to Lstat
				if path == startPath || *recursive {
					// Indicate that files in this dir are candidates.
					return nil
				}
				return filepath.SkipDir
			} else {
				// path is a regular file according to Lstat
				info, err = os.Stat(path)
				if err != nil {
					return err
				}
				if info.IsDir() {
					// path is a directory according to Stat

					// Symlinks whose targets are directories can result in
					// cycles, so just ignore it.
				} else {
					// path is a regular file according to Stat
					if matchRe.MatchString(basename) {
						candidates = append(candidates, path)
					}
				}
				return nil
			}
		})
	return candidates, err
}

func pick(candidatePaths []string) (string, error) {
	n := len(candidatePaths)
	if n == 0 {
		return "", errors.New("Nothing to pick from.")
	}
	return candidatePaths[rand.Int()%n], nil
}
