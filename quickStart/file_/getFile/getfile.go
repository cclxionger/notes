package getfile

import (
	"log"
	"os"
	"path/filepath"
)
const SourceDir = "sourcefile"
const Destfile = "destfile"
func GetFiles(dir string) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	filepath.Join()
	list := make([]string, 0)
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		fullName := filepath.Join(dir, f.Name())
		list = append(list, fullName)
	}
	return list
}
