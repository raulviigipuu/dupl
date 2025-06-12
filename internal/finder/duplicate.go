package finder

import (
	"os"
	"path/filepath"
)

// Find recursively walks through the directory tree starting at rootDirectoryPath,
// and returns a DuplicateFilesMap which groups files by name and size.
func Find(rootDirectoryPath string) (DuplicateFilesMap, error) {
	// Initialize the map to hold grouped files
	duplicateFilesMap := make(DuplicateFilesMap)

	// Start traversal
	// Note that structs, slices and channels are reference types, so I pointer goes to function
	err := checkDirectory(rootDirectoryPath, duplicateFilesMap)
	if err != nil {
		return nil, err
	}

	return duplicateFilesMap, nil
}

// checkDirectory visits all files in the given path and
// updates the provided DuplicateFilesMap with them.
func checkDirectory(currentPath string, duplicateFilesMap DuplicateFilesMap) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(currentPath, entry.Name())

		if entry.IsDir() {
			// Recurse into subdirectory
			if err := checkDirectory(fullPath, duplicateFilesMap); err != nil {
				return err
			}
		} else {
			info, err := entry.Info()
			if err != nil {
				return err
			}

			fileData := FileData{
				Name: info.Name(),
				Size: info.Size(),
			}

			duplicateFilesMap[fileData] = append(duplicateFilesMap[fileData], fullPath)
		}
	}

	return nil
}
