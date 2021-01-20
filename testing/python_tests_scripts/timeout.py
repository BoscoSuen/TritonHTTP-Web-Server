from socket import socket
import time
# Create connection to the server

s = socket()

s.connect(("localhost", 8080))

# Compose the message/HTTP request we want to send to the server

msgPart1 = b"GET /index.html HTTP/1.1\r\n"

msgPart2 = b"Host: Ha\r\n\r\n"

# Send out the request

s.sendall(msgPart1)

time.sleep(10)

s.sendall(msgPart2)

# Listen for response and print it out
recv = s.recv(4096)
print(recv)

# encoding bytes to a string
encode = 'utf-8'
output = str(recv, encode)

f = open("testing_result.txt", "a")
idx = output.index('\r\n\r\n')

f.write(f"# TESTING CASE: timeout: \n\n")
f.writelines(f">> Input: \n{msgPart1.decode(encode)}")
f.writelines(f">> Output: \n{output[:idx]}\n\n\n")

f.close()
s.close()
