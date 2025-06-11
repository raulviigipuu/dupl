package finder

import (
	"os"
	"path/filepath"
)

// FileData stores file name and size for comparison
type FileData struct {
	Name string
	Size int64
}

// Find searches for duplicate files in the specified root directory.
func Find(rootPath string) (map[FileData][]string, error) {
	filesMap := make(map[FileData][]string)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fd := FileData{Name: info.Name(), Size: info.Size()}
			filesMap[fd] = append(filesMap[fd], path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filesMap, nil
}
