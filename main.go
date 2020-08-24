package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

/*
Big File Finder, your BFF
To find all files under a path recursively which are above a minimum size limit.
*/

var (
	pathsAboveRedSizeMBLimit = flag.Int("minsize", 500, "minimum size to display in MBs")
	targetDir                = flag.String("path", "/tmp", "path to scan")
	version                  = flag.Bool("version", false, "show me version")
)

func readDir(dirname string) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

func sizeInHuman(size int64) (float64, string, string) {
	floatSize := float64(size)
	kbSize := floatSize / 1024
	if kbSize < 1.0 {
		return floatSize, fmt.Sprintf("%d bytes", size), "b"
	}
	mbSize := kbSize / 1024
	if mbSize < 1.0 {
		return kbSize, fmt.Sprintf("about %.2f KBs", kbSize), "kb"
	}
	gbSize := mbSize / 1024
	if gbSize < 1.0 {
		return mbSize, fmt.Sprintf("about %.2f MBs", mbSize), "mb"
	}
	return gbSize, fmt.Sprintf("about %.2f GBs", gbSize), "gb"
}

func calculateDirSize(dirpath string) (dirsize int64, err error) {
	err = os.Chdir(dirpath)
	if err != nil {
		return
	}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			dirsize += file.Size()
		}
	}
	return
}

func analyzeDir(pathToScan string) {
	sizeMBLimit := float64(*pathsAboveRedSizeMBLimit)
	if entries, err := readDir(pathToScan); err == nil {
		for _, entry := range entries {
			filepath := filepath.Join(pathToScan, entry.Name())
			size, humanSize, sizeType := sizeInHuman(entry.Size())
			isDir := entry.IsDir()
			if isDir {
				dirsize, err := calculateDirSize(filepath)
				if err != nil {
					log.Fatalf("not able to get dir size for: %s", filepath)
					continue
				} else if float64(dirsize) >= sizeMBLimit {
					analyzeDir(filepath)
				}
				continue
			}
			if sizeType == "b" || sizeType == "kb" {
				continue
			} else if sizeType == "mb" && size < sizeMBLimit {
				continue
			} else if sizeType == "gb" {
				mbSize := size * 1024
				if mbSize < sizeMBLimit {
					continue
				}
			}
			fmt.Printf("\n📁Path: %s\n\tName: %s\n\tSize: %s\n", filepath, entry.Name(), humanSize)
		}
	} else {
		log.Fatalf("error reading dir %s", pathToScan)
	}
}

func main() {
	flag.Parse()
	fmt.Println(`
 ====================================================
  ___ _        ___ _ _       ___ _         _
 | _ |_)__ _  | __(_) |___  | __(_)_ _  __| |___ _ _
 | _ \ / _| | | _|| | / -_) | _|| | ' \/ _| / -_) '_|
 |___/_\__, | |_| |_|_\___| |_| |_|_||_\__,_\___|_|
       |___/
  💣           📁            🕵️
 ====================================================
	`)
	if *version {
		fmt.Println(" Version: 0.0.1")
	} else {
		analyzeDir(*targetDir)
	}
	fmt.Println(`
 ====================================================
	`)
}
