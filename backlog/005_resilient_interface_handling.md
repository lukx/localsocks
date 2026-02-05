# Story: Resilient Network Interface Handling

As an operator, I want the application to be resilient to network interruptions, so that it can recover automatically when the selected interface goes down temporarily.

## Acceptance Criteria:

*   The application continuously monitors the status of the selected network interface.
*   If an established connection fails because the interface is down, the error is logged.
*   The application attempts to re-establish its capability to create connections on the interface.
*   Reconnection attempts use an exponential backoff strategy (e.g., wait 1s, 2s, 4s, 8s...).
*   Once the interface is available again, the application resumes normal operation and logs a confirmation.
