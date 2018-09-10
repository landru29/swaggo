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
	var vendorRegExp = regexp.MustCompile(`\/vendor\/`)
	err = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && goFileRegExp.MatchString(f.Name()) && !notHiddentDirRegExp.MatchString(path) && !vendorRegExp.MatchString(path) {
			fileList = append(fileList, path)
		}
		return nil
	})
	return
}

// ProcessProject is the main processor
func ProcessProject(searchDir string, host string, basePath string, schemes []string, verbose bool, filename string) (err error) {
	swag := NewSwagger(host, basePath, schemes, verbose)
	filenames, err := getFileList(searchDir)
	if err == nil {
		for _, filename := range filenames {
			if fileAnalyze, err := descriptor.ParseComments(filename); err == nil {
				GeneralInformations(&swag, &fileAnalyze, verbose)
				GetSubRoute(&swag, &fileAnalyze, verbose)
				GetDefinitions(&swag, &fileAnalyze, verbose)
			}
		}
		CompileSubRoutes(&swag, verbose)
		for _, filename := range filenames {
			if fileAnalyze, err := descriptor.ParseComments(filename); err == nil {
				Route(&swag, &fileAnalyze, verbose)
			}
		}
	}

	err = swag.Save(filename)
	return
}
