# OSI Model

Open System Interconnection Model. Introduced in 1984.

1. Physical Layer
2. Data Link Layer
3. Network Layer
4. Transport Layer
5. Session Layer
6. Presentation Layer
7. Application Layer

Each layer have bunch of protocols to work correctly.

## Basic Idea using image

![OSI Model Simplified](../Images/osi-image-1.png?raw=true "OSI Model In a Nutshell")

## Application Layer (Layer - 07)

- Any kind of application which create request or recieve response. Mostly browser, skype, game app etc. In case browser, it can create a GET request to server.

GET / 10.0.0.3:80
Http header, cookies, content-type etc.

It will convert this request to bunch of string with numbers. And send to next layer.

File Transfer - FTP
Email - SMTP
Web Surfing - HTTP/S
Virtual Terminal - Telnet

Protocols: HTTP, HTTPS, FTP, FMTP, TELNET, DCHP, NFS, IRC, POP3, NNTP.

## Presentation Layer (Layer - 06)

- Translation (String to Binary)
- Data Compression (For size reduction)
- Encrypt/Decrypt (if necessary)

So, if we use https or TLS, then this layer will encrypt that bunch of bytes otherwise it well pass those bytes to next layer.

Protocols: SSL

## Session Layer (Layer - 05)

- Session Management - Establish session and tag it to data.
- Authentication - Who you are?
- Authorization - Do you have permission to access this/that file?

So, if the request is stateful then session layer will add some bits to original bytes to ensure that this bytes belong to this session, so that it can be stateful. And obviously, session id will an encrypted thing.

## Transport Layer (Layer - 04 Segment)

- Segmentation
- Flow Control
- Error Control

Segmentation: Breaking bytes of data into managable small piece and give sequence number, call it as Segment. Also, assign port number to every segment, source port(kinda auto-generated) & Destination Port.

Flow Control: Server can send data 100 Mbps speed but mobile can receive in 10 Mbps, so this layer maintain this flow control.

Error Control: If one data unit is missing to other side, then it will resend that segment. Checksum(group of extra bit) is used to identify corrupted segment.

Protocols: TCP, UDP

TCP (Transmission Control Protocol) - Connection Oriented Transmission - Need ACK. (WWW, Email, FTP)
UDP (User Datagram Protocol) - Connetionless Transmission - No ACK. (VOIP Calls, Streaming, DNS Call, TFTP)

## Network Layer (Layer - 03 Packet)

- Logical Addressing
- Routing
- Path Determination

Logical Addressing: Added IP with each segment, source & destination IP. Call it as packet here.

Routing: Added mask to that packet.

Path Determination: From client to server, there are multiple ways to send/recieve data. So, layer-3 devices use OSPF (Open Shortest Path First), BGP (Border Gateway Protocol), IS-IS (Intermediate System - Int. Sys.) to find best possible shortest path to send data.

- Russian Doll Effect.

## Data Link Layer (Layer - 02 Frame)

- Physical Addressing (Adding MAC address)

- Break it down each packet and add destination & source mac address to each packet, call it as frame.

## Physical Layer (Layer - 01)

- Convert bits (010101010111000010100100100) to signal.

- Convert to electic signal in case of copper media.

- Convert to light signal in case of optical fiber.

- Convert bits to radio signals in case of air.

### Thanks

This file is based on 2 videos, here is the link:

- [Hussein Nasser](https://www.youtube.com/watch?v=7IS7gigunyI)
- [TechTerms](https://www.youtube.com/watch?v=vv4y_uOneC0)
