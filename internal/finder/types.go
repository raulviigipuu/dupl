package finder

// FileData is a lightweight signature of a file (name + size).
type FileData struct {
	Name string
	Size int64
}

// DuplicateFilesMap groups files by their name and size signature.
type DuplicateFilesMap map[FileData][]string
