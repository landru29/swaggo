package swagger

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"

    "github.com/spf13/viper"
)

// Save save the swagger
func (data Swagger) Save() (err error) {
    jsonStr, err := json.Marshal(data)
    if err != nil {
        return
    }

    fileHandle, err := os.Create(viper.GetString("output"))
    if err != nil {
        return
    }
    writer := bufio.NewWriter(fileHandle)
    defer fileHandle.Close()

    fmt.Fprintln(writer, string(jsonStr))
    writer.Flush()

    return
}
