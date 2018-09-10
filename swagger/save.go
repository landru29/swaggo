package swagger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Save save the swagger
func (data Swagger) Save(filename string) (err error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return
	}

	fileHandle, err := os.Create(filename)
	if err != nil {
		return
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, string(jsonStr))
	writer.Flush()

	return
}
