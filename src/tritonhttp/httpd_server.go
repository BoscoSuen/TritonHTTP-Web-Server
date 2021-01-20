package tritonhttp

import (
    "log"
    "net"
)

/** 
    Initialize the tritonhttp server by populating HttpServer structure
**/
func NewHttpdServer(port, docRoot, mimePath string) (*HttpServer, error) {
    // Initialize mimeMap for server to refer
    var httpServer HttpServer
    httpServer.DocRoot = docRoot
    httpServer.ServerPort = port
    httpServer.MIMEPath = mimePath
    httpServer.MIMEMap = make(map[string]string)
    err := ParseMIME(mimePath, httpServer.MIMEMap)
    if err != nil {
        log.Panicln(err)
    }

    // Return pointer to HttpServer
    return &httpServer, err
}

/** 
    Start the tritonhttp server
**/
func (hs *HttpServer) Start() (err error) {
    // Start listening to the server port
    listen, err := net.Listen("tcp", hs.ServerPort)
    if err != nil {
        log.Panicln(err)
    }
    log.Println("Start listing the server.")
    log.Println("Port: ", hs.ServerPort)
    log.Println("Doc Root: ", hs.DocRoot)

    // Accept connection from client
    for {
        conn, err := listen.Accept()
        // log.Println(conn.RemoteAddr())
        if err != nil {
            log.Panicln(err)
        }
        go hs.handleConnection(conn)
    }
    // Spawn a go routine to handle request

}

