package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// FileData stores file name and size for comparison
type FileData struct {
	Name string
	Size int64
}

func main() {
	// Define command-line flags
	rootPath := flag.String("path", ".", "Specify the root directory path to search for duplicate files (default is current directory)")
	pathAlias := flag.String("p", "", "Alias for -path")
	help := flag.Bool("help", false, "Display help information")
	helpAlias := flag.Bool("h", false, "Display help information (short)")

	// Parse the flags
	flag.Parse()

	// Check if help was requested with either -help or -h
	if *help || *helpAlias {
		flag.Usage()
		return
	}

	// Check if the -p alias was used for -path
	if *pathAlias != "" {
		*rootPath = *pathAlias
	}

	// Print the root path being used and a hint for more usage info
	fmt.Printf("Searching for duplicate files in: %s\n", *rootPath)
	fmt.Println("Use -h for more options.")

	// Use the root directory path from the command-line flag
	duplicates := findDuplicateFiles(*rootPath)

	// Print out the duplicates, if found
	for _, files := range duplicates {
		if len(files) > 1 {
			fmt.Println("Duplicates found:")
			for _, file := range files {
				fmt.Println(file)
			}
		}
	}
}

// findDuplicateFiles traverses the given rootPath to find and return duplicate files
func findDuplicateFiles(rootPath string) map[FileData][]string {
	filesMap := make(map[FileData][]string)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileData := FileData{Name: info.Name(), Size: info.Size()}
			filesMap[fileData] = append(filesMap[fileData], path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through files: %v\n", err)
		return nil
	}

	return filesMap
}
