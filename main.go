package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/landru29/swaggo/parser"
)

/* Gros comments
sur deux lignes */

// ceci est un test
func getFileList(searchDir string) (fileList []string, err error) {
	fileList = []string{}
	var goFileRegExp = regexp.MustCompile(`\.go$`)
	var notHiddentDirRegExp = regexp.MustCompile(`\/\.\w+|^\.\w+`)
	err = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && goFileRegExp.MatchString(f.Name()) && !notHiddentDirRegExp.MatchString(path) {
			fileList = append(fileList, path)
		}
		return nil
	})
	return
}

// et un autre
func main() {
	filenames, err := getFileList(".")
	if err == nil {
		for _, filename := range filenames {
			fmt.Printf("############ %s ############\n", filename)
			comments, _ := parser.ParseComments(filename)
			for _, comments := range comments {
				fmt.Println(comments)
			}
		}
	}
}
