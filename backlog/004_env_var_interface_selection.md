# Story: Interface Selection via Environment Variable

As an operator, I want to configure the desired network interface using an environment variable so that I can easily change the setting without recompiling the code.

## Acceptance Criteria:

*   The application reads the `LOCALSOCKS_INTERFACE` environment variable at startup.
*   If the variable is set, the application uses the specified interface name (e.g., "en0", "eth1") for upstream binding.
*   If the variable is not set, the application may default to the primary system interface but should log a clear warning that it's not bound to a specific interface.
*   If the specified interface does not exist, the application should log a fatal error and exit.
