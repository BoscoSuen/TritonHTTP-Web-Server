package tritonhttp

import (
    "bufio"
    "io"
    "log"
    "os"
    "strings"
)

/** 
    Load and parse the mime.types file 
**/
func ParseMIME(MIMEPath string, mimeMap map[string]string) (err error) {
    file, err := os.Open(MIMEPath)
    if err != nil {
        return err
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)
        extension := parts[0]
        extensionType := parts[1]
        mimeMap[extension] = extensionType
    }

    return err
}

/**
    Create file with content.
**/
func CreateFile(path string, content string) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        file, _ := os.Create(path)
        _, fileErr :=  io.WriteString(file, content)
        if fileErr != nil {
            log.Panicln(fileErr)
        }
	}
}