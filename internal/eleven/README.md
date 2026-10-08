# Eleven Layer

This package is the product layer that sits above the Sanaei/3x-ui core.

The Eleven layer must remain isolated from Xray-specific implementation details wherever possible.

## Initial domains

- identity: administrators, roles, permissions
- service: users, groups, templates, subscriptions
- nodes: node registry and health
- audit: immutable operational events
- integration: versioned API for Eleven Store
