package main

import (
	"flag"
	"fmt"
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
	listDir                  = flag.Bool("dir", false, "list big size directories also")
	version                  = flag.Bool("version", false, "show me version")
	debug                    = flag.Bool("debug", false, "show extra information")
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

func changeSizeToMB(size float64, sizeType string) float64 {
	if sizeType == "b" {
		return (size / (1024 * 1024))
	} else if sizeType == "kb" {
		return (size / (1024))
	} else if sizeType == "mb" {
		return size
	} else if sizeType == "gb" {
		return size * 1024
	}
	return size
}

func calculateDirSize(dirPath string) (dirSize int64, fullDirSize float64, err error) {
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			dirSize += info.Size()
		}
		return nil
	})

	size, _, sizeType := sizeInHuman(dirSize)
	fullDirSize = changeSizeToMB(size, sizeType)
	return
}

func analyzeSubDir(dirPath string, dirName string, sizeMBLimit float64) {
	_, dirSize, err := calculateDirSize(dirPath)
	if err != nil && *debug == true {
		log.Printf("not able to get dir size for: %s", dirPath)
		return
	} else if float64(dirSize) >= sizeMBLimit {
		if *listDir {
			fmt.Printf("\n📁Path: %s\n\tName: %s\n\tSize: %.2f MB\n", dirPath, dirName, dirSize)
		}
		analyzeDir(dirPath)
	}
}

func analyzeDir(pathToScan string) {
	sizeMBLimit := float64(*pathsAboveRedSizeMBLimit)
	if entries, err := readDir(pathToScan); err == nil {
		for _, entry := range entries {
			filepath := filepath.Join(pathToScan, entry.Name())
			size, humanSize, sizeType := sizeInHuman(entry.Size())

			if entry.IsDir() {
				analyzeSubDir(filepath, entry.Name(), sizeMBLimit)
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
			fmt.Printf("\n📄Path: %s\n\tName: %s\n\tSize: %s\n", filepath, entry.Name(), humanSize)
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
