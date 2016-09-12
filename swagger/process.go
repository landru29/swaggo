package swagger

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/landru29/swaggo/descriptor"
)

// Build the list of files to scan
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

// ProcessProject is the main processor
func ProcessProject(searchDir string, host string, basePath string, schemes []string) (err error) {
	swag := NewSwagger(host, basePath, schemes)
	filenames, err := getFileList(searchDir)
	if err == nil {
		for _, filename := range filenames {
			fileAnalyze, _ := descriptor.ParseComments(filename)
			GeneralInformations(&fileAnalyze, &swag)
		}
		for _, filename := range filenames {
			fileAnalyze, _ := descriptor.ParseComments(filename)
			SubRoute(&fileAnalyze, &swag)
		}
		for _, filename := range filenames {
			fileAnalyze, _ := descriptor.ParseComments(filename)
			Route(&fileAnalyze, &swag)
		}
	}

	err = swag.Save()
	return
}