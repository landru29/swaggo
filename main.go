package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
)

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

func parseCommentLines(filename string) (comments []string, err error) {
    comments = []string{}
    dat, err := ioutil.ReadFile(filename)
    contentFile := string(dat)
    inlineCommentRegExp := regexp.MustCompile(`\/\/\s?(.*)`)
    inlineComment := inlineCommentRegExp.FindAllString(contentFile, -1)
    for _, comment := range inlineComment {
        filtered := inlineCommentRegExp.FindStringSubmatch(comment)
        comments = append(comments, filtered[1])
    }
    blockCommentRegExp := regexp.MustCompile(`(?sg)\/\*(\*(?!\/)|[^*])*\*\/`)

    return
}

// et un autre
func main() {
    filenames, err := getFileList(".")
    if err == nil {
        for _, filename := range filenames {
            comments, _ := parseCommentLines(filename)
            for _, comments := range comments {
                fmt.Println(comments)
            }
        }
    }
}
