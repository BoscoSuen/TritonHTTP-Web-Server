from socket import socket
# import sys

port = 8088

# Create connection to the server
s = socket()
s.connect(("localhost", port))

# Compose the message/HTTP request we want to send to the server

msgPart1 = b"GET HTTP/1.1\r\n\r\n"

# Send out the request

s.sendall(msgPart1)

# Listen for response and print it out
f = open("response.output", "a")

recv = s.recv(4096)
print(recv)

# encoding bytes to a string
encode = 'utf-8'
output = str(recv, encode)

idx = output.index('\r\n\r\n')

f.write(f"# TESTING: empty header path \n\n")
f.writelines(f">> Input: \n{msgPart1.decode(encode)}")
f.writelines(f">> Output: \n{output[:idx]}\n\n\n")

s.close()
