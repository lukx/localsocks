# Story: Graceful Handling of Interface Disappearance

As an operator, I want the application to react gracefully if the configured network interface disappears completely, so that the application doesn't crash and provides clear feedback.

## Acceptance Criteria:

*   The application detects if the configured network interface is no longer present on the system (e.g., a USB Wi-Fi dongle is unplugged or a VPN is disconnected).
*   When the interface disappears, the application logs a critical error message stating the interface is gone.
*   All active proxy connections are terminated cleanly.
*   The application stops listening for new client connections.
*   The application exits with a non-zero status code to indicate failure.
