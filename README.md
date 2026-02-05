# LocalSOCKS Proxy

LocalSOCKS is a simple SOCKS5 proxy written in Go that allows you to bind client application traffic to a specific local network interface.

This is useful for applications that do not support explicit network interface binding, allowing you to control their network path through their proxy settings.

## Features

- **SOCKS5 Proxy:** A fully functional SOCKS5 proxy server.
- **Specific Interface Binding:** Binds all upstream connections to a chosen network interface.
- **Dynamic Port Allocation:** Automatically finds an available port to run on, starting from port 10800.
- **Resilient:** Gracefully handles temporary network interface outages with an exponential backoff retry mechanism.
- **Environment-based Configuration:** The network interface is easily configured via an environment variable.

## Building

To build the application, you need to have Go installed (version 1.18 or later).

Clone the repository and then run the following command from the project root:

```bash
go build
```

This will create a `localsocks` executable in the project directory.

## Running

To run the proxy, you must set the `LOCALSOCKS_INTERFACE` environment variable to the name of the network interface you want to use.

For example, on macOS, your Wi-Fi interface might be `en0`. On Linux, it might be `eth0` or `wlan0`. You can find your interface names by running `ifconfig` (on macOS) or `ip addr` (on Linux).

```bash
# Example for macOS
LOCALSOCKS_INTERFACE="en0" ./localsocks

# Example for Linux
LOCALSOCKS_INTERFACE="eth0" ./localsocks
```

If the `LOCALSOCKS_INTERFACE` is not set, the proxy will run but will use the system's default interface for outbound connections.

Upon startup, the application will log the address and port it is listening on, for example:
`SOCKS proxy listening on 127.0.0.1:10800`

## Usage

To use the proxy, you need to configure your client application's SOCKS proxy settings to point to the address and port logged by the LocalSOCKS application.

- **Host/Address:** `127.0.0.1`
- **Port:** The port number logged on startup (e.g., `10800`)

For example, in Firefox's network settings, you would configure it like this:

1. Go to `Settings` -> `General` -> `Network Settings` -> `Settings...`
2. Select `Manual proxy configuration`.
3. Set the `SOCKS Host` to `127.0.0.1`.
4. Set the `Port` to the one provided by the `localsocks` application.
5. Make sure `SOCKS v5` is selected.

All traffic from that application will now be routed through the LocalSOCKS proxy and will exit your system via the network interface you specified.
