package tritonhttp

type HttpServer	struct {
    ServerPort	string
    DocRoot		string
    MIMEPath	string
    MIMEMap		map[string]string
}

type HttpResponseHeader struct {
    // Add any fields required for the response here
    Server          string
    Last_Modified   string
    Content_Type    string
    Content_Length  int64
    Status          string
    Code            int
    FilePath        string
    CloseConn       bool
}

type HttpRequestHeader struct {
    // Add any fields required for the request here
    Header      map[string]string
    Status      int
    RequestDir  string
    CloseConn   bool
    Valid       bool
}