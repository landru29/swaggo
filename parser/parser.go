package parser

import (
    "io/ioutil"
    "regexp"
    "strings"
)

// ParseComments parse all the comments
func ParseComments(filename string) (comments []string, err error) {
    comments = []string{}
    dat, err := ioutil.ReadFile(filename)
    contentFile := string(dat)

    // Inline comments
    inlineCommentRegExp := regexp.MustCompile(`\/\/\s?(.*)`)
    inlineComment := inlineCommentRegExp.FindAllString(contentFile, -1)
    for _, comment := range inlineComment {
        filtered := inlineCommentRegExp.FindStringSubmatch(comment)
        comments = append(comments, filtered[1])
    }

    // block comments
    blockCommentRegExp := regexp.MustCompile(`\/\*(\*([^\/])|[^*])*\*\/`)
    blockComment := blockCommentRegExp.FindAllStringSubmatch(contentFile, -1)
    for _, comment := range blockComment {
        commentStr := strings.Replace(comment[0], "/"+"*", "", -1)
        commentStr = strings.Replace(commentStr, "*"+"/", "", -1)
        lines := strings.Split(commentStr, "\n")
        for _, line := range lines {
            comments = append(comments, line)
        }
    }

    return
}
