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
func ProcessProject(searchDir string, host string, basePath string, schemes []string, verbose bool) (err error) {
	swag := NewSwagger(host, basePath, schemes, verbose)
	filenames, err := getFileList(searchDir)
	if err == nil {
		for _, filename := range filenames {
			if fileAnalyze, err := descriptor.ParseComments(filename); err == nil {
				swag.GeneralInformations(&fileAnalyze, verbose)
				swag.GetSubRoute(&fileAnalyze, verbose)
			}
		}
		swag.CompileSubRoutes(verbose)
		for _, filename := range filenames {
			if fileAnalyze, err := descriptor.ParseComments(filename); err == nil {
				swag.Route(&fileAnalyze, verbose)
			}
		}
	}

	err = swag.Save()
	return
}
