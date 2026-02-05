# Story: Dynamic Port Allocation

As an operator, I want the application to automatically find and bind to an available port so that I don't have to manually configure it and avoid conflicts.

## Acceptance Criteria:

*   The application defines a default starting port (e.g., 10800).
*   On startup, the application attempts to bind to the current port.
*   If the port is already in use, it increments the port number by 1 and retries.
*   This process continues until a free port is successfully bound.
*   The application logs the port it has successfully bound to.
