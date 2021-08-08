# OSI Model

Open System Interconnection Model.

1. Physical Layer
2. Data Link Layer
3. Network Layer
4. Transport Layer
5. Session Layer
6. Presentation Layer
7. Application Layer

## Application Layer (Layer - 07)

- Any kind of application which create request or recieve response. Mostly browser, skype, game app etc. In case browser, it can create a GET request to server.

GET / 10.0.0.3:80
Http header, cookies, content-type etc.

It will convert this request to bunch of string, eventually bunch of bytes. And send to next layer.

## Presentation Layer (Layer - 06)

- Encrypt if necessary.

So, if we use https or TLS, then this layer will encrypt that bunch of bytes otherwise it well pass those bytes to next layer.

## Session Layer (Layer - 05)

- Establish session and tag it to data.

So, if the request is stateful then session layer will add some bits to original bytes to ensure that this bytes belong to this session, so that it can be stateful. And obviously, session id will an encrypted thing.

## Transport Layer (Layer - 04)

- Breaking bytes of data into managable small piece and obviously maintain sequence, call it as Segment.

- Assign port number to every segment. Source port(kinda auto-generated) & Destination Port.

## Network Layer (Layer - 03)

- Assing(added) IP with each segment, source IP & destination IP. Call it as packet here.

- Each segment has multiple packet. Here, we don't check for error.

- Russian Doll Effect.

## Data Link Layer (Layer - 02)

- Break it down each packet and add destination & source mac address to each packet, call it as frame.

## Physical Layer (Layer - 01)

- Convert every bits (010101010111000010100100100) which is frame actually into electic signal.
