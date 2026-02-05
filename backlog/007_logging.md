# Story: Application Logging

As an operator, I want clear and concise logging so that I can understand the status and health of the application.

## Acceptance Criteria:

*   Upon successfully starting the proxy and binding to a port and interface, a confirmation message is logged to standard output (e.g., "SOCKS proxy listening on 127.0.0.1:10800, bound to upstream interface en0").
*   After startup, only warnings, errors, and critical failure messages are logged.
*   Routine connection handling and data transfer are not logged.
*   Logs are sent to standard error.
