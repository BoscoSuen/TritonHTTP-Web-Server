
package tritonhttp

import (
    "bufio"
    "log"
    "net"
    "path/filepath"
    "strings"
    "strconv"
    "os"
)

func (hs *HttpServer) handleBadRequest(conn net.Conn) {
    log.Println("Bad request: 400 Bad Request")
    var responseHeader HttpResponseHeader
    responseHeader.Server = "Server: Go-Triton-Server-1.0"

    responseHeader.Status = "400 Bad Request"
    responseHeader.Code = 400
    responseHeader.CloseConn = true
    responseHeader.FilePath = hs.DocRoot + "/bad_request.html"
    content := "<html>\n\t<head>\n\t\t<title>400 Bad Request</title>\n\t</head>\n\t<body>\n\t\t<h1>400 Bad Request</h1>\n\t</body>\n</html>\n"
    CreateFile(responseHeader.FilePath, content)
    info, _ := os.Stat(responseHeader.FilePath)
    responseHeader.Content_Length = info.Size()
    responseHeader.Content_Type = hs.MIMEMap[".html"]
    hs.sendResponse(responseHeader, conn)
}

func (hs *HttpServer) handleFileNotFoundRequest(requestHeader *HttpRequestHeader, conn net.Conn) {
    log.Println("File Not Found: 404") 
    var responseHeader HttpResponseHeader
    responseHeader.Server = "Server: Go-Triton-Server-1.0"

    responseHeader.Status = "404 Not Found"
    responseHeader.Code = 404
    responseHeader.FilePath = hs.DocRoot + "/file_not_found.html"
    content := "<html>\n\t<head>\n\t\t<title>400 Bad Request</title>\n\t</head>\n\t<body>\n\t\t<h1>404 Not Found</h1>\n\t</body>\n</html>\n"
    CreateFile(responseHeader.FilePath, content)
    info, _ := os.Stat(responseHeader.FilePath)
    responseHeader.Content_Length = info.Size()
    responseHeader.Content_Type = hs.MIMEMap[".html"]
    hs.sendResponse(responseHeader, conn)
}

func (hs *HttpServer) handleResponse(requestHeader *HttpRequestHeader, conn net.Conn) (result bool) {
    log.Println("Handle response")
    var responseHeader HttpResponseHeader
    responseHeader.Server = "Server: Go-Triton-Server-1.0"

    if requestHeader.CloseConn {
        responseHeader.CloseConn = true
    }

    filePath := hs.DocRoot + requestHeader.RequestDir
    // log.Println("filepath: ", filePath)
    info, err := os.Stat(filePath)
    if err != nil {
        log.Println("file not found")  
        hs.handleFileNotFoundRequest(requestHeader, conn)
        return
    }

    // Check if escape doc root
    absDocRoot, _ := filepath.Abs(hs.DocRoot)
    absFilePath, _ := filepath.Abs(filePath)
    // log.Println("absd: ", absDocRoot)
    // log.Println("absf: ", absFilePath)
    if !strings.Contains(absFilePath, absDocRoot) {
        log.Println("Escape the Doc Root!")
        hs.handleFileNotFoundRequest(requestHeader, conn)
        return
    }

    responseHeader.Last_Modified = info.ModTime().String()
    responseHeader.Content_Length = info.Size()
    tokens := strings.Split(filePath, ".")
    extension := "." + tokens[len(tokens) - 1]
    // log.Println("extension", extension)
    if hs.MIMEMap[extension] != "" {
        responseHeader.Content_Type = hs.MIMEMap[extension]
    } else {
        log.Println("default MIME type")
        responseHeader.Content_Type = "application/octet-stream"
    }

    responseHeader.Status = "200 OK"
    responseHeader.Code = 200
    responseHeader.FilePath = filePath
    hs.sendResponse(responseHeader, conn)

    return true
}

/**
* Sample response:
* HTTP/1.1 200 OK
* Server: Go-Triton-Server-1.0 <CRLF>
* Last-Modified: Fri, 02 Oct 2020 14:37:09 -0700 <CRLF>
* Content-Length: 307<CRLF>
* Content-Type: text/html <CRLF>
* <CRLF>
* <optional body>
*/
func (hs *HttpServer) sendResponse(responseHeader HttpResponseHeader, conn net.Conn) {
    // Send headers
    log.Println("Send response: ", responseHeader.Status)
    CRLF := "\r\n"
    response := "HTTP/1.1 " + responseHeader.Status + CRLF
    response += "Server: Go-Triton-Server-1.0" + CRLF
    if responseHeader.Code == 200 {
        response += "Last-Modified: " + responseHeader.Last_Modified + CRLF
    }
    response += "Content-Length: " + strconv.FormatInt(responseHeader.Content_Length, 10) + CRLF
    response += "Content-Type: " + responseHeader.Content_Type + CRLF
    if responseHeader.CloseConn {
        response += "Connection: Close" + CRLF
    }
    // log.Println("responseHeader.Content_Type: ", responseHeader.Content_Type)
    response += CRLF

    _, err := conn.Write([]byte(response))
    if err != nil {
        log.Println("Write Error: ", err)
    }


    // Send file if required

    file, err := os.Open(responseHeader.FilePath)

    reader := bufio.NewReader(file)
    writer := bufio.NewWriter(conn)

    defer writer.Flush()

    _, writeErr := writer.ReadFrom(reader)
    if err != nil {
        log.Panicln(writeErr)
    }
    
    // Hint - Use the bufio package to write response
}
