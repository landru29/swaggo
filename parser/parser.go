package parser

import (
    "io/ioutil"
    "regexp"
    "sort"
    "strings"
)

//FileAnalyze is the analyse of a source file
type FileAnalyze struct {
    Package  string
    Comments []string
}

type commentStruct struct {
    Value string
    Pos   int
}

type byIndex []commentStruct

func (a byIndex) Len() int {
    return len(a)
}

func (a byIndex) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a byIndex) Less(i, j int) bool {
    return a[i].Pos < a[j].Pos
}

// ParseComments parse all the comments
func ParseComments(filename string) (analyse FileAnalyze, err error) {
    comments := []commentStruct{}
    analyse.Comments = []string{}
    dat, err := ioutil.ReadFile(filename)
    contentFile := string(dat)

    // Inline comments
    inlinecommentStructRegExp := regexp.MustCompile(`\/\/\s?(.*)`)
    inlinecommentStruct := inlinecommentStructRegExp.FindAllString(contentFile, -1)
    inlinecommentStructIndex := inlinecommentStructRegExp.FindAllStringIndex(contentFile, -1)
    for key, comment := range inlinecommentStruct {
        filtered := inlinecommentStructRegExp.FindStringSubmatch(comment)
        comments = append(comments, commentStruct{filtered[1], inlinecommentStructIndex[key][0]})
    }

    // block comments
    blockcommentStructRegExp := regexp.MustCompile(`\/\*(\*([^\/])|[^*])*\*\/`)
    blockcommentStruct := blockcommentStructRegExp.FindAllStringSubmatch(contentFile, -1)
    blockcommentStructIndex := blockcommentStructRegExp.FindAllStringIndex(contentFile, -1)
    for key, comment := range blockcommentStruct {
        commentStr := strings.Replace(comment[0], "/"+"*", "", -1)
        commentStr = strings.Replace(commentStr, "*"+"/", "", -1)
        lines := strings.Split(commentStr, "\n")
        for index, line := range lines {
            comments = append(comments, commentStruct{line, blockcommentStructIndex[key][0] + index})
        }
    }

    // package
    packageRegExp := regexp.MustCompile(`package\s*([^\n\s]*)`)
    packageFind := packageRegExp.FindStringSubmatch(contentFile)
    if len(packageFind) > 1 {
        analyse.Package = packageFind[1]
    }

    // compile all
    sort.Sort(byIndex(comments))
    for _, comment := range comments {
        analyse.Comments = append(analyse.Comments, comment.Value)
    }

    return
}

// MatchField try to get a field comment
func MatchField(commentStr string, name string) (values []string, ok bool) {
    findRegExp := regexp.MustCompile(`^\s*@` + name + `(.*)`)
    params := findRegExp.FindStringSubmatch(commentStr)
    ok = (len(params) == 2)
    if ok {
        values = strings.Fields(params[1])
    }
    return
}

// GetFields search for all fields in a list of comments
func GetFields(commentStrList []string, name string) (values [][]string) {
    values = [][]string{}
    for _, commentStr := range commentStrList {
        lineValues, ok := MatchField(commentStr, name)
        if ok {
            values = append(values, lineValues)
        }
    }
    return
}

// GetField search for a field in a list of comments
func GetField(commentStrList []string, name string) (values []string, ok bool) {
    result := GetFields(commentStrList, name)
    ok = false
    if len(result) > 0 {
        ok = true
        values = result[0]
    }
    return
}
