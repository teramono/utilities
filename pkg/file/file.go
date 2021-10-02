package file

import (
	"os"
	"path/filepath"
)

// OpenOrCreateFile ...
func OpenOrCreateFile(path string, isAppend bool) (*os.File, error) {
	var append int

	if isAppend {
		append = os.O_APPEND
	} else {
		append = 0
	}

	// Had to pass 0777 here. https://stackoverflow.com/a/58403214/3984876
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return nil, err
	}

	// SEC: No point giving other non-grp users permissions.
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|append, 0660)
}

// CreateFolderIfNotExist ...
func CreateFolderIfNotExist(path string) error {
	// Had to pass 0777 here. https://stackoverflow.com/a/58403214/3984876
	return os.MkdirAll(path, 0777)
}

// ReadFile ...
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// Canonicalize ...
func Canonicalize(path string, canonicalPWD string) string {
	// TODO: SEC: Harden!
	return ""
}
