package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	log.SetFlags(0)

	recurs := flag.Bool("r", false, "Also pick from nested directory contents.")
	flag.Parse()

	candidatePaths := search(".", *recurs)
	n := len(candidatePaths)
	if n == 0 {
		log.Fatal("Nothing to pick from.")
	}
	rand.Seed(time.Now().UnixNano())
	pick := candidatePaths[rand.Int()%n]
	log.Println("Picked:", pick)

	var launcher string
	switch runtime.GOOS {
	case "darwin":
		launcher = "open"
	case "windows":
		launcher = "start"
	default:
		launcher = "xdg-open"
	}
	err := exec.Command(launcher, pick).Start()
	if err != nil {
		log.Fatal(err)
	}
}

func search(startPath string, recurs bool) []string {
	candidatePaths := []string{}
	visitor := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			if !recurs && path != startPath {
				return filepath.SkipDir
			}
		} else {
			if info.Mode()&os.ModeSymlink > 0 {
				info, err = os.Stat(path)
				if err != nil {
					log.Fatal(err)
				}
			}
			if !info.IsDir() {
				candidatePaths = append(candidatePaths, path)
			}
		}
		return nil
	}
	filepath.Walk(startPath, visitor)
	return candidatePaths
}
