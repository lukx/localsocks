# Story: Basic SOCKS Proxy Server

As a user, I want a basic SOCKS5 proxy server so that I can route traffic through it.

## Acceptance Criteria:

*   The application starts a SOCKS5 server.
*   The server listens for incoming connections on a local port.
*   The server can handle the SOCKS5 handshake.
*   The server can establish a connection to a remote destination requested by the client.
*   The server relays data between the client and the destination.
