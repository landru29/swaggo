package descriptor

import (
	"regexp"
	"strings"
)

// MatchField try to get a field comment
func MatchField(commentStr string, name string) (values []string, ok bool) {
	findRegExp := regexp.MustCompile(`(?i)^\s*@` + name + `(.*)`)
	params := findRegExp.FindStringSubmatch(commentStr)
	ok = (len(params) == 2)
	if ok {
		splitRegExp := regexp.MustCompile(`"[^"]+"|[^\s"]+`)
		splits := splitRegExp.FindAllStringSubmatch(params[1], -1)
		for _, p := range splits {
			values = append(values, p[0])
		}
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

// DescID decode a comment line like "hello world [id]"
func DescID(comment []string) (id string, description string, comments []string, ok bool) {
	ok = false
	if len(comment) < 1 {
		return
	}

	resStr := comment[len(comment)-1]
	resRegExp := regexp.MustCompile(`^\[([^\]]+)\]$`)
	resMatch := resRegExp.FindStringSubmatch(resStr)
	if len(resMatch) < 2 {
		return
	}
	id = resMatch[1]
	comments = comment[:len(comment)-1]
	description = strings.Join(comments, " ")
	ok = true
	return
}
