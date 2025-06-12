package main

import (
	"flag"
	"fmt"

	"github.com/raulviigipuu/dupl/internal/finder"
	"github.com/raulviigipuu/dupl/internal/logx"
)

var Version = "dev" // default if not overridden at build time

func main() {

	logx.Init(nil)
	logx.Info("Starting dupl...")

	// Define command-line flags
	rootPath := flag.String("path", ".", "Specify root directory (default: current dir)")
	pathAlias := flag.String("p", "", "Alias for -path")
	help := flag.Bool("help", false, "Show help")
	helpAlias := flag.Bool("h", false, "Show help (short)")
	versionFlag := flag.Bool("v", false, "Show version and exit")

	flag.Parse()

	// Help
	if *help || *helpAlias {
		printHelp()
		return
	}

	// Version
	if *versionFlag {
		logx.Info(fmt.Sprintf("dupl version: %s", Version))
		return
	}

	if *pathAlias != "" {
		*rootPath = *pathAlias
	}

	logx.Info(fmt.Sprintf("Searching for duplicate files in: %s", *rootPath))

	duplicates, err := finder.Find(*rootPath)
	if err != nil {
		logx.FatalErr(err)
	}

	for _, files := range duplicates {
		if len(files) > 1 {
			logx.Info("Duplicates found:")
			for _, file := range files {
				logx.Info(file)
			}
		}
	}
}

func printHelp() {
	fmt.Println("dupl — Fast duplicate file finder")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  dupl [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -path string     Root directory to scan (default: \".\")")
	fmt.Println("  -p string        Alias for -path")
	fmt.Println("  -v               Show version and exit")
	fmt.Println("  -h, -help        Show this help and exit")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  dupl -p /data/photos")
	fmt.Println("  dupl -path ./backup")
	fmt.Println("  dupl -v")
}
