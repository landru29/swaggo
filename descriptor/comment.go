package descriptor

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

//FileAnalyze is the analyse of a source file
type FileAnalyze struct {
	Filename      string
	FileComments  []string
	BlockComments [][]string
}

type commentStruct struct {
	Value string
	Line  int
}

type byIndex []commentStruct

func (a byIndex) Len() int {
	return len(a)
}

func (a byIndex) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byIndex) Less(i, j int) bool {
	return a[i].Line < a[j].Line
}

// get the line of a comment
func getLine(globalPos int, indexes [][]int) int {
	if len(indexes) > 0 {
		if globalPos <= indexes[0][0] {
			return 1
		}
		for index := 1; index < len(indexes); index++ {
			if (globalPos <= indexes[index][0]) && (globalPos >= indexes[index-1][1]) {
				return index + 1
			}
		}
	}
	return 0
}

func gatherComments(comments []commentStruct) (result [][]string) {
	result = [][]string{}
	cursor := -10
	var block []string
	sort.Sort(byIndex(comments))
	for _, comment := range comments {
		if comment.Line != cursor+1 {
			if cursor != -10 {
				result = append(result, block)
			}
			block = []string{}
		}
		block = append(block, comment.Value)
		cursor = comment.Line
	}
	if len(block) > 0 {
		result = append(result, block)
	}
	return
}

// ParseComments parse all the comments
func ParseComments(filename string) (analyse FileAnalyze, err error) {
	analyse.Filename = filename
	//fmt.Printf("############ %s ############", filename)
	comments := []commentStruct{}
	analyse.FileComments = []string{}
	dat, err := ioutil.ReadFile(filename)
	contentFile := string(dat)

	// get position of carriage return
	carriageReturnRegExp := regexp.MustCompile(`\n`)
	carriageReturnIndex := carriageReturnRegExp.FindAllStringIndex(contentFile, -1)

	// Inline comments
	inlinecommentStructRegExp := regexp.MustCompile(`\/\/\s?(.*)`)
	inlinecommentStruct := inlinecommentStructRegExp.FindAllString(contentFile, -1)
	inlinecommentStructIndex := inlinecommentStructRegExp.FindAllStringIndex(contentFile, -1)
	for key, comment := range inlinecommentStruct {
		filtered := inlinecommentStructRegExp.FindStringSubmatch(comment)
		comments = append(
			comments,
			commentStruct{
				filtered[1],
				getLine(inlinecommentStructIndex[key][0], carriageReturnIndex),
			})
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
			comments = append(
				comments,
				commentStruct{
					line,
					getLine(blockcommentStructIndex[key][0], carriageReturnIndex) + index,
				})
		}
	}

	// compile all
	sort.Sort(byIndex(comments))
	for _, comment := range comments {
		//fmt.Printf("# %d - %s\n", comment.Line, comment.Value)
		analyse.FileComments = append(analyse.FileComments, comment.Value)
	}

	analyse.BlockComments = gatherComments(comments)

	return
}
