package tritonhttp

import (
    "log"
    "net"
    "io"
    "strings"
    "time"
)

/* 
For a connection, keep handling requests until 
    1. a timeout occurs or
    2. client closes connection or
    3. client sends a bad request
*/
func (hs *HttpServer) handleConnection(conn net.Conn) {
    // Start a loop for reading requests continuously
    // Set a timeout for read operation
    // Read from the connection socket into a buffer
    // ____Validate the request lines that were read____
    // Handle any complete requests
    // Update any ongoing requests
    // If reusing read buffer, truncate it before next read

    log.Println("New connection accepted.")
    defer log.Println("Connection closed.")

    timeout := 5 * time.Second 
    // var headers []HttpRequestHeader     // Multiple request headers in one TCP connection.
    CRLF := "\r\n"
    requestEndDelim := CRLF + CRLF
    remaining := ""                     // Whole HTTP request.

    for {
        // log.Println("addr", conn.RemoteAddr())
        conn.SetReadDeadline(time.Now().Add(timeout))  

        buf := make([]byte, 1024 * 8)
        size, err := conn.Read(buf)
        // println("size: ",size) 
        if err != nil || size == 0 {
            if remaining != "" {
                log.Println("Timeout with partial header")
                hs.handleBadRequest(conn)    // partial header but timeout
                conn.Close()
                return
            } else if err == io.EOF {
                conn.Close()
                return
            }
        }
        // Get all request headers.
        data := string(buf[:size])
        remaining += data
        if (strings.Contains(remaining, requestEndDelim)) {
            index := strings.LastIndex(remaining, requestEndDelim)
            requests := strings.Split(remaining[:index], requestEndDelim)
            for _, request := range requests {
                log.Println("Deal with a new request")

                var header HttpRequestHeader
                header.Header = make(map[string]string)
                requestParts := strings.Split(request, CRLF)
                decodeInitialLine(&header, requestParts[0])
                if header.Status == 400 {
                    hs.handleBadRequest(conn)     // Bad request: invalid init line.
                    conn.Close()
                    return
                }
                for _, line := range requestParts[1:] {
                    if line == "" {
                        continue
                    }
                    decodeHeaderLine(&header, line)
                    
                    if header.Status == 400 {
                        break
                    }
                }
                if _, ok := header.Header["Host"]; ok && header.Status != 400 {
                    header.Status = 200
                } else {
                    header.Status = 400    // Host required.
                }
    
                if header.Header["Connection"] == "close" || header.Header["Connection"] == "Close" {
                    header.CloseConn = true   // Close connection
                }

                if header.Status == 200 {
                    hs.handleResponse(&header, conn)
                } else {
                    hs.handleBadRequest(conn)
                    conn.Close()
                    return
                }
                if header.CloseConn {
                    log.Println("Connection close set by header")
                    conn.Close()       // Close connection
                    return
                }
            }
            remaining = remaining[index + 4:]       // Trim reading buffer.
            // log.Println("remaining: ", remaining)
        }
    }
}

/**
* Sample initial line:
* GET /index.html HTTP/1.1
*/
func decodeInitialLine(header *HttpRequestHeader, initLine string) {
    initTokens := strings.Fields(initLine)
    if len(initTokens) != 3 {
        header.Status = 400
        return
    }
    if (initTokens[0] != "GET" || initTokens[2] != "HTTP/1.1") {
        // Only accept GET and HTTP/1.1 protocol.
        header.Status = 400
        return
    }
    dir := initTokens[1]
    // If dir is "/", it will redirect to the index.html page
    if dir[0] == '/' {
        if strings.LastIndex(dir, ".") == -1 {
            dir += "/"      // end with dir
        }

        if dir[len(dir) - 1] == '/' {
            header.RequestDir = dir + "index.html"
        } else {
            header.RequestDir = dir
        }
        header.Status = 200
    } else {
        header.Status = 400     // dir doesn't start with '/'
    }
 }

 /**
 * Sample line:
 * Host: MyHost
 */
 func decodeHeaderLine(header *HttpRequestHeader, line string) {
    idx := strings.Index(line, ":")
    if idx <= 0 {
        header.Status = 400     // Invalid header format.
        return
    }
    key := strings.TrimSpace(line[:idx])
    val := strings.TrimSpace(line[idx + 1:])
    if val == "" {      // empty header value
        header.Status = 400
        return
    }
    header.Header[key] = val
 }
