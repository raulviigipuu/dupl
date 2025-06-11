package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/raulviigipuu/dupl/internal/finder"
	"github.com/raulviigipuu/dupl/internal/logx"
)

func main() {

	logx.Init(nil)
	logx.Info("Starting dupl")

	// Define command-line flags
	rootPath := flag.String("path", ".", "Specify root directory (default: current dir)")
	pathAlias := flag.String("p", "", "Alias for -path")
	help := flag.Bool("help", false, "Show help")
	helpAlias := flag.Bool("h", false, "Show help (short)")

	flag.Parse()

	if *help || *helpAlias {
		flag.Usage()
		return
	}

	if *pathAlias != "" {
		*rootPath = *pathAlias
	}

	fmt.Printf("Searching for duplicate files in: %s\n", *rootPath)

	duplicates, err := finder.Find(*rootPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, files := range duplicates {
		if len(files) > 1 {
			fmt.Println("Duplicates found:")
			for _, file := range files {
				fmt.Println(file)
			}
		}
	}
}
