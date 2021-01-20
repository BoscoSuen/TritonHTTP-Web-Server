printf 'GET /index.html HTTP/1.1\r\nHost: MyHost\r\n\r\n GET /index1.html HTTP/1.1\r\n
Host: MyHost\r\n\r\n GET /subdir1/index.html HTTP/1.1\r\nHost: MyHost\r\n\r\n' > request.input

cat request.input | nc localhost 8088 >> response.output