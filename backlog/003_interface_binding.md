# Story: Upstream Interface Binding

As a user, I want the proxy's outbound connections to be bound to a specific network interface, so that I can control the network path of my applications.

## Acceptance Criteria:

*   The application identifies a specific network interface to use for all upstream connections.
*   When the proxy establishes a connection to a remote destination on behalf of a client, that connection is made through the specified network interface.
*   System-level routing tables and default gateways are ignored in favor of this explicit binding.
