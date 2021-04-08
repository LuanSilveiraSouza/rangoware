package explorer

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

func MapFiles(dir string) []string {
	var files []string
	var root string

	if runtime.GOOS == "windows" {
		root = os.Getenv("USERPROFILE")
	} else {
		root = os.Getenv("HOME")
	}

	if dir == "" {
		root += "/Downloads/Test"
	} else {
		root += dir
	}

	error := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		files = append(files, path)

		return nil
	})

	if error != nil {
		panic(error)
	}

	return files
}
