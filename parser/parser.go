package parser

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

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
func ParseComments(filename string) (commentStr []string, err error) {
	comments := []commentStruct{}
	commentStr = []string{}
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

	// compile all
	sort.Sort(byIndex(comments))
	for _, comment := range comments {
		commentStr = append(commentStr, comment.Value)
	}

	return
}
