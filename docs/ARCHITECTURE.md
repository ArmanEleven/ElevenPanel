# Eleven Panel Architecture

## Principle

Keep the Sanaei/3x-ui core stable. Add Eleven features around it instead of scattering product logic through every existing handler.

## Target structure

```
Eleven Panel
├── Sanaei / 3x-ui Core
│   ├── Xray
│   ├── Inbounds
│   ├── Clients
│   ├── Hosts
│   └── Existing API
│
└── Eleven Layer
    ├── Identity
    │   ├── Admins
    │   ├── Roles
    │   └── Permissions
    ├── Service
    │   ├── Users
    │   ├── Groups
    │   ├── Templates
    │   └── Subscriptions
    ├── Nodes
    │   ├── Node registry
    │   ├── Health checks
    │   └── Node groups
    ├── Operations
    │   ├── Statistics
    │   ├── Activity logs
    │   └── Bulk operations
    └── Integration
        └── Eleven Store API
```

## Database direction

The first schema proposal uses separate Eleven tables rather than modifying existing Sanaei tables unnecessarily.

Core entities:

- eleven_admins
- eleven_roles
- eleven_permissions
- eleven_admin_roles
- eleven_role_permissions
- eleven_users
- eleven_groups
- eleven_templates
- eleven_nodes
- eleven_node_groups
- eleven_subscriptions
- eleven_hwids
- eleven_activity_logs

Foreign keys and indexes will be added only after the actual Sanaei schema is inspected.

## API direction

Eleven Store will communicate with Eleven Panel through a versioned API:

`/api/v1/eleven/...`

The API must support authentication, idempotency for provisioning requests, structured errors, audit logging, and safe retries.

## Migration

Migration is a release requirement:

Sanaei/3x-ui data -> Eleven migration adapter -> Eleven schema

The original data must never be destructively modified during migration.
