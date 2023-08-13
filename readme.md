# Simple websocket client-server communication example 
- This program basically demonstrate the basic chat application

# Basic information about websocket 
- WebSockets run on top of the TCP (Transmission Control Protocol) network protocol. 
- WebSockets provide a full-duplex communication channel over a single TCP connection, allowing for real-time data exchange between a client and a server.

- Here's a brief overview of how WebSockets use the TCP protocol:

- TCP is a reliable,connection-oriented protocol - mechanisms for establishing and maintaining a reliable connection between the sender and receiver.

- WebSocket Protocol: WebSockets build on top of TCP to provide a standardized way for web browsers and servers 
- to establish a long-lived, bidirectional communication channel. 
- Unlike traditional HTTP, where the client initiates a request and the server responds, 
- ==>WebSockets allow both the client and server to send messages to each other at any time without the overhead of re-establishing connections.

- Handshake: The WebSocket connection starts with an HTTP-based handshake, during which the client sends an upgrade request to the server,
- requesting to switch from HTTP to the WebSocket protocol.
- If the server supports WebSocket, it responds with a successful handshake, and the connection is upgraded to WebSocket.

- Data Exchange: Once the WebSocket connection is established, both the client and server can send messages in either direction. 
- These messages can be in the form of text or binary data. 
- The messages are framed within the WebSocket protocol, which includes framing information to delineate individual messages.

- Efficiency: WebSockets are more efficient for real-time communication compared to traditional HTTP polling 
- or long polling, as they eliminate the need for repeated connection establishment and HTTP headers in each message.

- Overall, WebSockets leverage the reliability and established infrastructure of the TCP protocol
- while providing an efficient and low-latency means of bidirectional communication over the web.